// Test du routage URL - parseURL et buildURL
const LICENSES = [
  { id: "PPL", label: "PPL" },
  { id: "LAPL", label: "LAPL" },
  { id: "CPL", label: "CPL" },
  { id: "ATPL", label: "ATPL" },
  { id: "IR", label: "IR" },
];

const CATEGORIES = [
  { id: "meteorology" },
  { id: "navigation" },
  { id: "airlaw" },
  { id: "aircraft_general" },
  { id: "performance" },
  { id: "human_performance" },
  { id: "operational_procedures" },
  { id: "communications" },
  { id: "principles_of_flight" },
  { id: "flight_planning" },
  { id: "instrumentation" },
  { id: "emergency" },
  { id: "mass_and_balance" },
  { id: "radio_procedure" },
];

// ═══ parseURL (copie exacte du code dans app.js) ═══
function parseURL(pathname) {
  const parts = pathname.replace(/\/+$/, "").split("/").filter(Boolean);
  if (parts.length === 0) return { view: "home", params: {} };
  const first = parts[0];
  const simpleRoutes = {
    "login": "login-form",
    "dashboard": "dashboard",
    "quiz": "quiz",
    "history": "history",
    "stats": "stats",
    "recommendations": "recommendations",
  };
  if (simpleRoutes[first]) {
    return { view: simpleRoutes[first], params: {} };
  }
  if (first === "questions") {
    if (parts.length === 1) return { view: "questions", params: {} };
    if (parts.length === 2) return { view: "questions-license", params: { license: parts[1] } };
    if (parts.length === 3 && parts[1] === "detail") return { view: "question-detail", params: { questionId: parts[2] } };
    if (parts.length === 3) return { view: "questions-category", params: { license: parts[1], category: parts[2] } };
    if (parts.length === 4) return { view: "question-detail", params: { license: parts[1], category: parts[2], questionId: parts[3] } };
  }
  if (first === "lessons") {
    if (parts.length === 1) return { view: "lessons", params: {} };
    if (parts.length === 2) return { view: "lessons-license", params: { license: parts[1] } };
    if (parts.length === 3 && parts[1] === "detail") return { view: "lesson-detail", params: { lessonId: parts[2] } };
    if (parts.length === 3) return { view: "lessons-category", params: { license: parts[1], category: parts[2] } };
    if (parts.length === 4) return { view: "lesson-detail", params: { license: parts[1], category: parts[2], lessonId: parts[3] } };
  }
  return { view: "home", params: {} };
}

// ═══ buildURL (copie exacte du code dans app.js) ═══
function buildURL(view, params = {}) {
  switch (view) {
    case "home":
    case "login":        return "/";
    case "login-form":   return "/login";
    case "dashboard":    return "/dashboard";
    case "questions":    return "/questions";
    case "questions-license":
      return params.license ? "/questions/" + params.license : "/questions";
    case "questions-category":
      if (params.license && params.category) return "/questions/" + params.license + "/" + params.category;
      if (params.license) return "/questions/" + params.license;
      return "/questions";
    case "question-detail":
      if (params.license && params.category && params.questionId)
        return "/questions/" + params.license + "/" + params.category + "/" + params.questionId;
      if (params.questionId) return "/questions/detail/" + params.questionId;
      return "/questions";
    case "quiz":         return "/quiz";
    case "lessons":      return "/lessons";
    case "lessons-license":
      return params.license ? "/lessons/" + params.license : "/lessons";
    case "lessons-category":
      if (params.license && params.category) return "/lessons/" + params.license + "/" + params.category;
      if (params.license) return "/lessons/" + params.license;
      return "/lessons";
    case "lesson-detail":
      if (params.license && params.category && params.lessonId)
        return "/lessons/" + params.license + "/" + params.category + "/" + params.lessonId;
      if (params.lessonId) return "/lessons/detail/" + params.lessonId;
      return "/lessons";
    case "history":      return "/history";
    case "stats":        return "/stats";
    case "recommendations": return "/recommendations";
    default:             return "/dashboard";
  }
}

// ═══ TESTS ═══
let failures = 0;
let count = 0;

function eq(a, b, label) {
  if (a !== b) {
    console.log("  FAIL " + label + ": " + a + " !== " + b);
    failures++;
  }
  count++;
}

function testParse(url, expectedView, expectedParams) {
  const r = parseURL(url);
  eq(r.view, expectedView, url + " view");
  for (const k of Object.keys(expectedParams)) {
    eq(r.params[k] || "", expectedParams[k] || "", url + " params." + k);
  }
  // Pas de paramètres inattendus
  for (const k of Object.keys(r.params || {})) {
    if (!(k in expectedParams)) {
      eq("unexpected:" + k, "", url + " unexpected param " + k);
    }
  }
}

// Test de build → parse roundtrip
function testRoundtrip(view, params) {
  const url = buildURL(view, params);
  const r = parseURL(url);
  eq(r.view, view, "RT " + view + " -> " + url + " view");
  if (params.license) eq(r.params.license, params.license, "RT " + view + " license");
  if (params.category) eq(r.params.category, params.category, "RT " + view + " category");
  if (params.questionId) eq(r.params.questionId, params.questionId, "RT " + view + " questionId");
  if (params.lessonId) eq(r.params.lessonId, params.lessonId, "RT " + view + " lessonId");
}

console.log("=== PARSE URL TESTS ===");

// Routes simples
testParse("/", "home", {});
testParse("/login", "login-form", {});
testParse("/dashboard", "dashboard", {});
testParse("/quiz", "quiz", {});
testParse("/history", "history", {});
testParse("/stats", "stats", {});
testParse("/recommendations", "recommendations", {});

// Questions
testParse("/questions", "questions", {});
testParse("/questions/PPL", "questions-license", { license: "PPL" });
testParse("/questions/ppl", "questions-license", { license: "ppl" });
testParse("/questions/CPL", "questions-license", { license: "CPL" });
testParse("/questions/atpl", "questions-license", { license: "atpl" });
testParse("/questions/IR", "questions-license", { license: "IR" });
testParse("/questions/LAPL", "questions-license", { license: "LAPL" });
testParse("/questions/PPL/meteorology", "questions-category", { license: "PPL", category: "meteorology" });
testParse("/questions/PPL/meteorology/abc123", "question-detail", { license: "PPL", category: "meteorology", questionId: "abc123" });
testParse("/questions/detail/abc123", "question-detail", { questionId: "abc123" });

// Lessons
testParse("/lessons", "lessons", {});
testParse("/lessons/CPL", "lessons-license", { license: "CPL" });
testParse("/lessons/cpl", "lessons-license", { license: "cpl" });
testParse("/lessons/ATPL", "lessons-license", { license: "ATPL" });
testParse("/lessons/CPL/navigation", "lessons-category", { license: "CPL", category: "navigation" });
testParse("/lessons/CPL/navigation/xyz", "lesson-detail", { license: "CPL", category: "navigation", lessonId: "xyz" });
testParse("/lessons/detail/xyz", "lesson-detail", { lessonId: "xyz" });

// Fallback
testParse("/nimportequoi", "home", {});
testParse("/foo/bar/baz", "home", {});

console.log("");
console.log("=== ROUNDTRIP TESTS ===");

// buildURL → parseURL doit donner les mêmes valeurs
testRoundtrip("home", {});
testRoundtrip("login-form", {});
testRoundtrip("dashboard", {});
testRoundtrip("quiz", {});
testRoundtrip("history", {});
testRoundtrip("stats", {});
testRoundtrip("recommendations", {});
testRoundtrip("questions", {});
testRoundtrip("lessons", {});

// Toutes les licences
for (const l of LICENSES) {
  testRoundtrip("questions-license", { license: l.id });
  testRoundtrip("lessons-license", { license: l.id });
  // Toutes les catégories
  for (const c of CATEGORIES) {
    testRoundtrip("questions-category", { license: l.id, category: c.id });
    testRoundtrip("lessons-category", { license: l.id, category: c.id });
  }
}

// Détails
testRoundtrip("question-detail", { license: "PPL", category: "meteorology", questionId: "q1" });
testRoundtrip("question-detail", { questionId: "q1" });
testRoundtrip("lesson-detail", { license: "CPL", category: "navigation", lessonId: "l1" });
testRoundtrip("lesson-detail", { lessonId: "l1" });

console.log("");
console.log("=== RESULT ===");
console.log(count + " assertions, " + failures + " failures");
if (failures === 0) console.log("PASS");
else process.exit(1);