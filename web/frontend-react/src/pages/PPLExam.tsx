import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Navbar from "@/components/Navbar";
import { apiClient } from "@/api/client";
import { usePPLExamStore, type PPLPaper } from "@/store/pplExamStore";
import { useLearningLocale } from "@/utils/useLearningLocale";
import { renderMarkdown } from "@/utils/renderMarkdown";

export default function PPLExam() {
  const s = usePPLExamStore();
  const { text, choices } = useLearningLocale();
  const [papers, setPapers] = useState<PPLPaper[]>([]);
  const [loadError, setLoadError] = useState(false);
  const [index, setIndex] = useState(0);
  const [now, setNow] = useState(Date.now());
  const [confirmFinish, setConfirmFinish] = useState(false);
  useEffect(() => { apiClient.get<{ papers: PPLPaper[] }>("/api/ppl/exam-papers")
    .then(r => setPapers(r.data.papers)).catch(() => setLoadError(true)); }, []);
  useEffect(() => { const timer = setInterval(() => setNow(Date.now()), 500); return () => clearInterval(timer); }, []);
  useEffect(() => { if (s.phase === "active" && now >= s.deadline) void s.finish(); }, [now, s.phase, s.deadline, s.finish]);
  useEffect(() => { setIndex(0); setConfirmFinish(false); }, [s.paper?.id, s.deadline]);
  const remaining = Math.max(0, Math.ceil((s.deadline - now) / 1000));
  const q = s.questions[index];
  const result = q && s.results[q.id];
  const score = s.paper && s.completed[s.paper.id];
  const allPassed = papers.length === 9 && papers.every(p => s.completed[p.id]?.passed);
  return <><Navbar /><main className="max-w-4xl mx-auto p-4 text-white">
    <Link to="/questions" className="text-blue-300">{text("← Questions", "← Questions")}</Link>
    <h1 className="text-2xl font-bold my-4">{text("Examen blanc PPL(A) — France", "PPL(A) practice exam — France")}</h1>
    <p className="text-slate-300 mb-4">{text("9 matières · 128 QCM · 75 % dans chaque matière. Questions d’entraînement originales ; programme encore en cours de complétion.", "9 subjects · 128 MCQs · 75% in each subject. Original practice questions; syllabus coverage is still being completed.")}</p>
    {s.error && <p role="alert" className="text-red-300 mb-4">{text(s.error, "The request failed. Your saved answers are retained; please retry.")}</p>}
    {s.phase === "idle" ? <>
      <p className="mb-4 text-slate-300">{text("Les corrections apparaissent après la remise de l’épreuve. Une question sans réponse vaut zéro point. Le temps continue si vous quittez la page. Les épreuves peuvent être passées séparément.", "Corrections appear after submitting a paper. Unanswered questions score zero. Time continues if you leave the page. Papers may be taken separately.")}</p>
      {loadError && <p role="alert">{text("Impossible de charger les matières. Rechargez la page.", "Unable to load papers. Reload the page.")}</p>}
      {!papers.length && !loadError && <p role="status">{text("Chargement…", "Loading…")}</p>}
      {allPassed && <p className="text-green-300 mb-4">{text("Les neuf dernières épreuves de cette session sont réussies.", "All nine latest papers in this session have been passed.")}</p>}
      <div className="grid sm:grid-cols-2 gap-3">{papers.map(p => <button disabled={s.busy} key={p.id} onClick={() => void s.start(p.id)} className="text-left p-5 rounded-xl bg-slate-800 border border-slate-600 disabled:opacity-50">
        <strong>{text(p.title_fr, p.title_en)}</strong><p>{p.count} {text("questions", "questions")} · {p.minutes} min</p>
        {s.completed[p.id] && <p className={s.completed[p.id].passed ? "text-green-300" : "text-amber-300"}>{s.completed[p.id].correct}/{p.count} — {s.completed[p.id].passed ? text("Réussi", "Passed") : text("À retravailler", "Review needed")}</p>}
      </button>)}</div>
    </> : <>
      <h2 className="text-xl font-semibold mb-3">{text(s.paper?.title_fr, s.paper?.title_en)}</h2>
      {s.phase === "active" && <p className="text-xl mb-4" role="timer">{text("Temps restant", "Time remaining")} : {Math.floor(remaining / 60)}:{String(remaining % 60).padStart(2, "0")}</p>}
      {s.phase === "grading" && <div className="p-4 bg-slate-800 rounded-xl"><p>{text("Épreuve terminée. Préparation des corrections…", "Paper closed. Preparing corrections…")}</p>
        {!s.busy && <button className="bg-blue-600 p-3 rounded mt-3" onClick={() => void s.finish()}>{text("Reprendre la correction", "Resume grading")}</button>}</div>}
      {s.phase === "finished" && score && <div className="p-4 bg-slate-800 rounded-xl mb-4">
        <p className="text-xl">{score.correct}/{score.total} — {score.passed ? text("Épreuve réussie", "Paper passed") : text("Épreuve à retravailler", "Paper needs review")}</p>
        <p>{text("Seuil : 75 %. Le résultat de cette matière ne valide pas les autres.", "Pass mark: 75%. This subject result does not pass the other subjects.")}</p>
        <button className="bg-blue-600 p-3 rounded mt-3" onClick={s.close}>{text("Choisir une autre épreuve", "Choose another paper")}</button>
      </div>}
      {s.phase !== "grading" && q && <>
        <nav aria-label={text("Questions de l’épreuve", "Paper questions")} className="flex gap-2 flex-wrap mb-4">{s.questions.map((item,i) => <button key={item.id} onClick={() => { setIndex(i); setConfirmFinish(false); }} aria-current={i === index ? "step" : undefined} aria-label={`${i+1} — ${s.answers[item.id] ? text("réponse choisie", "answered") : text("sans réponse", "unanswered")}`} className={`w-10 h-10 rounded border ${i === index ? "border-white bg-blue-700" : s.answers[item.id] ? "border-blue-400 bg-slate-700" : "border-slate-600"}`}>{i+1}</button>)}</nav>
        <article className="p-5 rounded-xl bg-slate-800">
          <p className="text-slate-300 mb-3">{index+1}/{s.questions.length}</p>
          <div dangerouslySetInnerHTML={{ __html: renderMarkdown(text(q.question_fr,q.question_en,q.question_es)) }} />
          <div className="grid gap-3 mt-5">{choices(q).map((choice,i) => { const letter=String.fromCharCode(65+i);return <button key={letter} disabled={s.phase !== "active" || now >= s.deadline} aria-pressed={s.answers[q.id] === letter} onClick={() => s.choose(q.id,letter)} className={`text-left rounded p-3 border ${s.answers[q.id] === letter ? "bg-blue-800 border-blue-300" : "bg-slate-700 border-slate-500"}`}>{letter}. {choice}</button>; })}</div>
          {s.phase === "finished" && result && <div className="mt-5 border-t border-slate-600 pt-4">
            <p className={result.correct ? "text-green-300" : "text-amber-300"}>{result.correct ? text("Réponse correcte", "Correct answer") : result.selected ? text("Réponse incorrecte", "Incorrect answer") : text("Sans réponse : 0 point", "Unanswered: 0 points")}</p>
            <p>{text("Bonne réponse : ", "Correct answer: ")}{choices(q)[q.options.findIndex(o => (typeof o === "string" ? o : o.text) === result.correct_answer)] || result.correct_answer}</p>
            <div dangerouslySetInnerHTML={{ __html: renderMarkdown(text(result.explanation_fr,result.explanation_en,result.explanation_es)) }} />
          </div>}
        </article>
        {s.phase === "active" && <div className="my-5">
          <p>{Object.keys(s.answers).length}/{s.questions.length} {text("réponses choisies", "questions answered")}</p>
          {confirmFinish ? <><p className="my-2">{text("Remettre définitivement cette épreuve ? Les réponses manquantes compteront zéro.", "Submit this paper now? Missing answers will score zero.")}</p><button className="bg-green-700 p-3 rounded mr-3" onClick={() => void s.finish()}>{text("Remettre", "Submit")}</button><button onClick={() => setConfirmFinish(false)}>{text("Continuer", "Continue")}</button></> : <button className="bg-green-700 p-3 rounded mt-3" onClick={() => setConfirmFinish(true)}>{text("Terminer l’épreuve", "Finish paper")}</button>}
        </div>}
      </>}
    </>}
  </main></>;
}
