"""Apply additive migrations and course imports in one guarded transaction."""
import subprocess
import sys
from pathlib import Path

root = Path(sys.argv[1]).resolve()
snapshot = """
CREATE FUNCTION pg_temp.personal_fingerprint() RETURNS jsonb LANGUAGE plpgsql AS $$
DECLARE tbl text; digest text; result jsonb := '{}'::jsonb;
BEGIN
  FOREACH tbl IN ARRAY ARRAY['students','student_question_history','user_progress',
    'user_gamification','user_quests','spaced_repetition','leaderboard_weekly'] LOOP
    EXECUTE format('SELECT md5(COALESCE(string_agg((to_jsonb(t)-''role'')::text, ''''
      ORDER BY (to_jsonb(t)-''role'')::text), '''')) FROM %I t', tbl) INTO digest;
    result := result || jsonb_build_object(tbl, digest);
  END LOOP;
  RETURN result;
END $$;
CREATE TEMP TABLE before_personal AS SELECT pg_temp.personal_fingerprint() AS digest;
CREATE TEMP TABLE before_question_ids AS SELECT id FROM questions;
"""
parts = ["BEGIN; SET LOCAL lock_timeout='15s'; SET LOCAL statement_timeout='120s';",
         "LOCK TABLE students,student_question_history,user_progress,user_gamification,"
         "user_quests,spaced_repetition,leaderboard_weekly IN SHARE MODE;", snapshot]
for name in ['005_student_roles.sql', '006_question_options_en.sql', '007_spanish_content.sql']:
    parts.append((root/'scripts/migrations'/name).read_text(encoding='utf-8'))
for program in ['ppl', 'lapl', 'cpl', 'atpl', 'ir']:
    source = (root/'scripts/seed'/f'{program}_content.sql').read_text(encoding='utf-8')
    parts.append('\n'.join(line for line in source.splitlines()
                           if line.strip() not in {'BEGIN;', 'COMMIT;'}))
parts.append("""
DO $$ BEGIN
  IF (SELECT digest FROM before_personal) IS DISTINCT FROM pg_temp.personal_fingerprint() THEN
    RAISE EXCEPTION 'Personal data changed; rolling back';
  END IF;
  IF EXISTS (SELECT id FROM before_question_ids EXCEPT SELECT id FROM questions) THEN
    RAISE EXCEPTION 'Existing question identifiers removed; rolling back';
  END IF;
END $$;
COMMIT;
""")
result = subprocess.run(['docker', 'exec', '-i', 'aeropath-postgres', 'psql',
                         '-X', '-v', 'ON_ERROR_STOP=1', '-U', 'aeropath', '-d', 'aeropath'],
                        input='\n'.join(parts), text=True, capture_output=True)
if result.returncode:
    # Keep SQL diagnostics local to the root-only backup folder supplied by the caller.
    Path(sys.argv[2]).write_text(result.stderr, encoding='utf-8')
    raise SystemExit('Content transaction failed; diagnostic saved privately, changes rolled back.')
print('Content imported; personal data and existing question IDs preserved.')
