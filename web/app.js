// ======================== CONFIGURATION ========================
const API_BASE = "";
let authToken = localStorage.getItem("aeropath_token");
let currentUser = null;

// ======================== GESTION DE LA LANGUE ========================
const SUPPORTED_LANGUAGES = [
  { code: "fr", label: "Français", flag: "�x!��x!�" },
  { code: "en", label: "English", flag: "�x!��x!�" },
  { code: "de", label: "Deutsch", flag: "�x!��x!�" },
  { code: "es", label: "Español", flag: "�x!��x!�" },
  { code: "it", label: "Italiano", flag: "�x!��x!�" },
];

let currentLang = localStorage.getItem("aeropath_lang") || "fr";

function setLanguage(langCode) {
  currentLang = langCode;
  localStorage.setItem("aeropath_lang", langCode);
  // Si l'utilisateur est connecté, on met à jour sa préférence sur le serveur
  if (authToken) {
    api("/api/me/lang", {
      method: "PATCH",
      body: JSON.stringify({ lang: langCode }),
    }).catch(() => {});
  }
  render();
}

function getCurrentLang() {
  return SUPPORTED_LANGUAGES.find(l => l.code === currentLang) || SUPPORTED_LANGUAGES[0];
}

function renderLanguageSelector() {
  const current = getCurrentLang();
  return `
    <div class="relative group">
      <button class="flex items-center gap-1.5 px-2.5 py-1.5 rounded-lg text-sm text-slate-300 hover:bg-slate-700 transition" title="Changer la langue">
        <span>${current.flag}</span>
        <span class="hidden sm:inline">${current.label}</span>
        <svg class="w-3 h-3 ml-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
        </svg>
      </button>
      <div class="absolute right-0 mt-1 w-44 bg-slate-800 border border-slate-700 rounded-xl shadow-xl opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-50">
        <div class="py-1">
          ${SUPPORTED_LANGUAGES.map(l => `
            <button onclick="setLanguage('${l.code}')" class="w-full flex items-center gap-3 px-4 py-2.5 text-sm text-slate-300 hover:bg-slate-700 hover:text-white transition ${l.code === currentLang ? 'bg-slate-700/50 text-white' : ''}">
              <span class="text-lg">${l.flag}</span>
              <span>${l.label}</span>
              ${l.code === currentLang ? '<span class="ml-auto text-blue-400">�S</span>' : ''}
            </button>
          `).join("")}
        </div>
      </div>
    </div>
  `;
}

// Helper pour afficher le texte dans la bonne langue
function t(fr, en) {
  return currentLang === "fr" ? (fr || en) : (en || fr);
}


// ======================== �0TAT DE L'APPLICATION ========================
const state = {
  view: "login",
  questions: [],
  currentQuestion: null,
  currentLicense: null,
  currentCategory: null,
  history: [],
  stats: null,
  recommendations: null,
  lessons: [],
  currentLesson: null,
  quizQuestions: [],
  quizIndex: 0,
  quizScore: 0,
  quizFeedback: null,
  offlineQueue: [],
};

// ======================== ENREGISTREMENT SW ========================
if ("serviceWorker" in navigator) {
  // Nettoyer les anciens caches et SW avant d'enregistrer le nouveau
  window.addEventListener("load", async () => {
    // Désenregistrer tous les anciens SW
    const registrations = await navigator.serviceWorker.getRegistrations();
    for (const reg of registrations) {
      await reg.unregister();
      console.log("[PWA] Ancien SW désenregistré");
    }

    // Vider tous les caches
    const cacheKeys = await caches.keys();
    await Promise.all(cacheKeys.map((key) => caches.delete(key)));
    console.log("[PWA] Caches vidés");

    // Enregistrer le nouveau SW
    navigator.serviceWorker
      .register("/sw.js")
      .then((reg) => {
        console.log("[PWA] Service Worker enregistré:", reg.scope);

        reg.addEventListener("updatefound", () => {
          const newWorker = reg.installing;
          newWorker.addEventListener("statechange", () => {
            if (newWorker.state === "installed" && navigator.serviceWorker.controller) {
              showToast("Nouvelle version disponible ! Rafraîchissez la page.");
            }
          });
        });
      })
      .catch((err) => console.error("[PWA] Erreur SW:", err));
  });
}

// ======================== API CLIENT ========================
async function api(path, options = {}) {
  const headers = { "Content-Type": "application/json", ...options.headers };
  if (authToken) {
    headers["Authorization"] = `Bearer ${authToken}`;
  }

  try {
    const response = await fetch(`${API_BASE}${path}`, {
      ...options,
      headers,
    });

    if (response.status === 401) {
      logout();
      return null;
    }

    // Si offline (503 du SW)
    if (response.status === 503) {
      showToast("Mode hors-ligne : les données peuvent être limitées");
      return { offline: true };
    }

    return await response.json();
  } catch (err) {
    // Hors-ligne : essayer le cache local
    showToast("Vous êtes hors-ligne. Les réponses seront synchronisées plus tard.");
    return { offline: true };
  }
}

// ======================== AUTHENTIFICATION ========================
async function login(email, password) {
  const data = await api("/auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
  if (data?.token) {
    authToken = data.token;
    localStorage.setItem("aeropath_token", authToken);
    await loadUser();
    navigate("dashboard");
    showToast("Connecté !");
  } else {
    showToast(data?.error || "Erreur de connexion", "error");
  }
}

async function register(email, password) {
  const data = await api("/auth/register", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
  if (data?.token) {
    authToken = data.token;
    localStorage.setItem("aeropath_token", authToken);
    await loadUser();
    navigate("dashboard");
    showToast("Compte créé !");
  } else {
    showToast(data?.error || "Erreur d'inscription", "error");
  }
}

function logout() {
  authToken = null;
  localStorage.removeItem("aeropath_token");
  currentUser = null;
  navigate("login");
  showToast("Déconnecté");
}

async function loadUser() {
  const data = await api("/api/me");
  if (data && !data.offline) {
    currentUser = data;
  }
}

// ======================== ROUTAGE URL ========================

function buildURL(view, params) {
  params = params || {};
  switch (view) {
    case "home": case "login": return "/";
    case "login-form": return "/login";
    case "dashboard": return "/dashboard";
    case "questions": return "/questions";
    case "questions-license": return params.license ? "/questions/" + params.license : "/questions";
    case "questions-category":
      if (params.license && params.category) return "/questions/" + params.license + "/" + params.category;
      if (params.license) return "/questions/" + params.license;
      return "/questions";
    case "question-detail":
      if (params.license && params.category && params.questionId) return "/questions/" + params.license + "/" + params.category + "/" + params.questionId;
      if (params.questionId) return "/questions/detail/" + params.questionId;
      return "/questions";
    case "quiz": return "/quiz";
    case "lessons": return "/lessons";
    case "lessons-license": return params.license ? "/lessons/" + params.license : "/lessons";
    case "lessons-category":
      if (params.license && params.category) return "/lessons/" + params.license + "/" + params.category;
      if (params.license) return "/lessons/" + params.license;
      return "/lessons";
    case "lesson-detail":
      if (params.license && params.category && params.lessonId) return "/lessons/" + params.license + "/" + params.category + "/" + params.lessonId;
      if (params.lessonId) return "/lessons/detail/" + params.lessonId;
      return "/lessons";
    case "history": return "/history";
    case "stats": return "/stats";
    case "recommendations": return "/recommendations";
    default: return "/dashboard";
  }
}

function parseURL(pathname) {
  var parts = pathname.replace(/\/+$/, "").split("/").filter(Boolean);
  if (!parts.length) return { view: "home", params: {} };
  var first = parts[0];
  var simple = { login: "login-form", dashboard: "dashboard", quiz: "quiz", history: "history", stats: "stats", recommendations: "recommendations" };
  if (simple[first]) return { view: simple[first], params: {} };
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

function applyParsedRoute(route) {
  state.view = route.view;
  if (route.params.license) state.currentLicense = route.params.license.toUpperCase();
  else state.currentLicense = null;
  if (route.params.category) state.currentCategory = route.params.category;
  else state.currentCategory = null;
  if (route.params.questionId) state.currentQuestion = route.params.questionId;
  if (route.params.lessonId) state.currentLesson = route.params.lessonId;
}

function onPopState() {
  var route = parseURL(window.location.pathname);
  applyParsedRoute(route);
  render();
}

// ======================== NAVIGATION ========================
function goBack() {
  history.back();
}

function navigate(view, params) {
  params = params || {};
  var url = buildURL(view, params);
  applyParsedRoute({ view: view, params: params });
  if (window.location.pathname !== url) {
    history.pushState(null, "", url);
  }
  render();
}

// ======================== RENDU ========================
function render() {
  const app = document.getElementById("app");

  if (!authToken) {
    if (state.view === "login" || state.view === "home") {
      app.innerHTML = renderHomePage();
      bindHomeEvents();
    } else if (state.view === "login-form") {
      app.innerHTML = renderLogin();
      bindLoginEvents();
    } else {
      history.replaceState(null, "", "/");
      state.view = "home";
      app.innerHTML = renderHomePage();
      bindHomeEvents();
    }
    return;
  }

  switch (state.view) {
    case "dashboard":
      app.innerHTML = renderDashboard();
      loadDashboardData();
      break;
    case "questions":
      app.innerHTML = renderQuestions();
      loadQuestions();
      break;
    case "questions-license":
      app.innerHTML = renderQuestionsByLicense();
      break;
    case "questions-category":
      app.innerHTML = renderQuestionsByCategory();
      break;
    case "quiz":
      app.innerHTML = renderQuiz();
      break;
    case "history":
      app.innerHTML = renderHistory();
      loadHistory();
      break;
    case "lessons":
      app.innerHTML = renderLessons();
      loadLessons();
      break;
    case "lessons-license":
      app.innerHTML = renderLessonsByLicense();
      break;
    case "lessons-category":
      app.innerHTML = renderLessonsByCategory();
      setTimeout(() => loadLessonsByCategory(state.currentLicense, state.currentCategory), 50);
      break;
    case "question-detail":
      app.innerHTML = renderQuestionDetail();
      setTimeout(() => loadQuestionDetail(state.currentQuestion), 50);
      break;
    case "lesson-detail":
      app.innerHTML = renderLessonDetail();
      setTimeout(() => loadLessonDetail(state.currentLesson), 50);
      break;
    case "stats":
      app.innerHTML = renderStats();
      loadStats();
      break;
    case "recommendations":
      app.innerHTML = renderRecommendations();
      loadRecommendations();
      break;
    default:
      app.innerHTML = renderDashboard();
      loadDashboardData();
  }
}

// ======================== PAGE D'ACCUEIL PUBLIQUE ========================
function renderHomePage() {
  return `
    <div class="min-h-screen bg-gradient-to-b from-slate-900 via-slate-800 to-slate-900">
      <!-- Navigation -->
      <nav class="bg-slate-900/80 backdrop-blur-sm border-b border-slate-700/50 sticky top-0 z-50">
        <div class="max-w-6xl mx-auto px-4">
          <div class="flex items-center justify-between h-16">
            <div class="flex items-center gap-3">
              <span class="text-3xl">�S�️</span>
              <span class="text-xl font-bold text-white">AeroPath</span>
            </div>
            <div class="flex items-center gap-3">
              ${renderLanguageSelector()}
              <button onclick="navigate('login-form')" class="px-5 py-2 text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition-all shadow-lg shadow-blue-600/25">
                Connexion
              </button>
              <button onclick="navigate('login-form'); setTimeout(() => document.getElementById('tab-register')?.click(), 100)" class="px-5 py-2 text-sm font-medium text-slate-300 border border-slate-600 hover:border-blue-500 hover:text-blue-400 rounded-lg transition-all">
                S'inscrire
              </button>
            </div>

          </div>
        </div>
      </nav>

      <!-- Hero Section -->
      <section class="relative overflow-hidden">
        <div class="absolute inset-0 bg-gradient-to-b from-blue-500/5 via-transparent to-transparent pointer-events-none"></div>
        <div class="max-w-6xl mx-auto px-4 py-20 md:py-32">
          <div class="text-center max-w-3xl mx-auto">
            <div class="text-7xl md:text-8xl mb-8 animate-float">�S�️</div>
            <h1 class="text-4xl md:text-6xl font-extrabold text-white mb-6 leading-tight">
              Votre formation aéronautique,
              <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-cyan-400">partout avec vous</span>
            </h1>
            <p class="text-lg md:text-xl text-slate-400 mb-10 max-w-2xl mx-auto leading-relaxed">
              Préparez vos licences PPL, LAPL, CPL, ATPL et IR avec des questions interactives,
              des leçons détaillées et un suivi personnalisé. Même sans connexion.
            </p>
            <div class="flex flex-col sm:flex-row gap-4 justify-center">
              <button onclick="navigate('login-form'); setTimeout(() => document.getElementById('tab-register')?.click(), 100)" class="px-8 py-4 bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-xl text-lg transition-all shadow-xl shadow-blue-600/30 hover:shadow-blue-600/50 hover:scale-105">
                �xa� Commencer gratuitement
              </button>
              <button onclick="navigate('login-form')" class="px-8 py-4 bg-slate-800 hover:bg-slate-700 text-slate-300 font-semibold rounded-xl text-lg border border-slate-700 transition-all hover:scale-105">
                �x�⬍�S�️ J'ai déjà un compte
              </button>
            </div>
          </div>
        </div>
      </section>

      <!-- Features Section -->
      <section class="py-16 md:py-24">
        <div class="max-w-6xl mx-auto px-4">
          <h2 class="text-3xl md:text-4xl font-bold text-white text-center mb-4">Pourquoi AeroPath ?</h2>
          <p class="text-slate-400 text-center mb-16 max-w-xl mx-auto">Une plateforme complète conçue par des pilotes, pour les pilotes</p>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div class="bg-slate-800/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50 hover:border-blue-500/30 transition-all group">
              <div class="w-14 h-14 bg-blue-500/10 rounded-xl flex items-center justify-center text-2xl mb-4 group-hover:bg-blue-500/20 transition-all">�xa</div>
              <h3 class="text-xl font-bold text-white mb-2">Questions illimitées</h3>
              <p class="text-slate-400 leading-relaxed">Des milliers de questions couvrant toutes les licences et tous les sujets, avec des explications détaillées pour chaque réponse.</p>
            </div>

            <div class="bg-slate-800/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50 hover:border-blue-500/30 transition-all group">
              <div class="w-14 h-14 bg-purple-500/10 rounded-xl flex items-center justify-center text-2xl mb-4 group-hover:bg-purple-500/20 transition-all">�x</div>
              <h3 class="text-xl font-bold text-white mb-2">Leçons interactives</h3>
              <p class="text-slate-400 leading-relaxed">Apprenez à votre rythme avec des leçons structurées, quiz intégrés et suivi de progression détaillé.</p>
            </div>

            <div class="bg-slate-800/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50 hover:border-blue-500/30 transition-all group">
              <div class="w-14 h-14 bg-green-500/10 rounded-xl flex items-center justify-center text-2xl mb-4 group-hover:bg-green-500/20 transition-all">�x`</div>
              <h3 class="text-xl font-bold text-white mb-2">Suivi personnalisé</h3>
              <p class="text-slate-400 leading-relaxed">Statistiques détaillées, recommandations intelligentes et répétition espacée pour optimiser votre apprentissage.</p>
            </div>

            <div class="bg-slate-800/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50 hover:border-blue-500/30 transition-all group">
              <div class="w-14 h-14 bg-amber-500/10 rounded-xl flex items-center justify-center text-2xl mb-4 group-hover:bg-amber-500/20 transition-all">�x�</div>
              <h3 class="text-xl font-bold text-white mb-2">Mode hors-ligne</h3>
              <p class="text-slate-400 leading-relaxed">Continuez à réviser même en vol ou dans les zones sans réseau. Synchronisation automatique à la reconnexion.</p>
            </div>

            <div class="bg-slate-800/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50 hover:border-blue-500/30 transition-all group">
              <div class="w-14 h-14 bg-red-500/10 rounded-xl flex items-center justify-center text-2xl mb-4 group-hover:bg-red-500/20 transition-all">�x}�</div>
              <h3 class="text-xl font-bold text-white mb-2">Recommandations IA</h3>
              <p class="text-slate-400 leading-relaxed">Notre moteur analyse vos résultats et vous suggère les sujets à réviser pour progresser plus vite.</p>
            </div>

            <div class="bg-slate-800/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50 hover:border-blue-500/30 transition-all group">
              <div class="w-14 h-14 bg-cyan-500/10 rounded-xl flex items-center justify-center text-2xl mb-4 group-hover:bg-cyan-500/20 transition-all">�x� </div>
              <h3 class="text-xl font-bold text-white mb-2">Toutes les licences</h3>
              <p class="text-slate-400 leading-relaxed">PPL, LAPL, CPL, ATPL, IR � préparez toutes vos certifications au même endroit.</p>
            </div>
          </div>
        </div>
      </section>

      <!-- Stats Section -->
      <section class="py-16 bg-slate-800/30">
        <div class="max-w-6xl mx-auto px-4">
          <div class="grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
            <div>
              <div class="text-4xl font-bold text-white mb-1">+5000</div>
              <div class="text-slate-400 text-sm">Questions</div>
            </div>
            <div>
              <div class="text-4xl font-bold text-white mb-1">+200</div>
              <div class="text-slate-400 text-sm">Leçons</div>
            </div>
            <div>
              <div class="text-4xl font-bold text-white mb-1">6</div>
              <div class="text-slate-400 text-sm">Licences</div>
            </div>
            <div>
              <div class="text-4xl font-bold text-white mb-1">100%</div>
              <div class="text-slate-400 text-sm">Hors-ligne</div>
            </div>
          </div>
        </div>
      </section>

      <!-- Licences Section -->
      <section class="py-16 md:py-24">
        <div class="max-w-6xl mx-auto px-4">
          <h2 class="text-3xl md:text-4xl font-bold text-white text-center mb-4">Licences disponibles</h2>
          <p class="text-slate-400 text-center mb-16 max-w-xl mx-auto">Du pilote privé au transport aérien, nous couvrons toutes les étapes</p>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div class="bg-gradient-to-br from-blue-900/20 to-slate-800 rounded-xl p-6 border border-blue-800/30">
              <div class="text-3xl mb-3">�x:�️</div>
              <h3 class="text-lg font-bold text-white mb-1">PPL / LAPL</h3>
              <p class="text-sm text-slate-400">Pilote Privé � La base de l'aviation</p>
            </div>
            <div class="bg-gradient-to-br from-purple-900/20 to-slate-800 rounded-xl p-6 border border-purple-800/30">
              <div class="text-3xl mb-3">�S�️</div>
              <h3 class="text-lg font-bold text-white mb-1">CPL</h3>
              <p class="text-sm text-slate-400">Pilote Professionnel � Faites de votre passion un métier</p>
            </div>
            <div class="bg-gradient-to-br from-amber-900/20 to-slate-800 rounded-xl p-6 border border-amber-800/30">
              <div class="text-3xl mb-3">�x:�</div>
              <h3 class="text-lg font-bold text-white mb-1">ATPL / IR</h3>
              <p class="text-sm text-slate-400">Transport Aérien & Vol aux Instruments � Le plus haut niveau</p>
            </div>
          </div>
        </div>
      </section>

      <!-- CTA Section -->
      <section class="py-16 md:py-24">
        <div class="max-w-4xl mx-auto px-4 text-center">
          <div class="bg-gradient-to-br from-blue-600/10 to-cyan-600/10 rounded-3xl p-10 md:p-16 border border-blue-500/20">
            <div class="text-6xl mb-6">�S�️</div>
            <h2 class="text-3xl md:text-4xl font-bold text-white mb-4">Prêt à décoller ?</h2>
            <p class="text-lg text-slate-400 mb-8 max-w-lg mx-auto">Rejoignez des centaines de pilotes qui préparent leur licence avec AeroPath</p>
            <button onclick="navigate('login-form'); setTimeout(() => document.getElementById('tab-register')?.click(), 100)" class="px-10 py-4 bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-xl text-lg transition-all shadow-xl shadow-blue-600/30 hover:shadow-blue-600/50 hover:scale-105">
              �xa� Créer mon compte gratuit
            </button>
          </div>
        </div>
      </section>

      <!-- Footer -->
      <footer class="border-t border-slate-800 py-8">
        <div class="max-w-6xl mx-auto px-4 text-center">
          <div class="flex items-center justify-center gap-2 mb-4">
            <span class="text-xl">�S�️</span>
            <span class="font-bold text-white">AeroPath</span>
          </div>
          <p class="text-slate-500 text-sm">Formation aéronautique pour pilotes � PPL, LAPL, CPL, ATPL, IR</p>
          <p class="text-slate-600 text-xs mt-2">© ${new Date().getFullYear()} AeroPath. Tous droits réservés.</p>
        </div>
      </footer>
    </div>
  `;
}

function bindHomeEvents() {
  // Animation au scroll
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.style.opacity = '1';
        entry.target.style.transform = 'translateY(0)';
      }
    });
  }, { threshold: 0.1 });

  document.querySelectorAll('section').forEach(section => {
    section.style.opacity = '0';
    section.style.transform = 'translateY(20px)';
    section.style.transition = 'all 0.6s ease-out';
    observer.observe(section);
  });
}

// ======================== PAGE DE CONNEXION ========================
function renderLogin() {
  return `
    <div class="min-h-screen bg-slate-900 flex items-center justify-center p-4">
      <div class="w-full max-w-md">
        <div class="text-center mb-8">
          <div class="text-5xl mb-4">�S�️</div>
          <h1 class="text-3xl font-bold text-white">AeroPath</h1>
          <p class="text-slate-400 mt-2">Formation aéronautique</p>
        </div>

        <div class="bg-slate-800 rounded-xl p-6 shadow-xl">
          <div class="flex mb-6">
            <button id="tab-login" class="flex-1 py-2 text-center font-medium rounded-l-lg bg-blue-600 text-white">Connexion</button>
            <button id="tab-register" class="flex-1 py-2 text-center font-medium rounded-r-lg bg-slate-700 text-slate-300">Inscription</button>
          </div>

          <form id="auth-form">
            <div class="mb-4">
              <label class="block text-sm text-slate-400 mb-1">Email</label>
              <input type="email" id="email" class="w-full bg-slate-700 text-white rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="pilote@example.com" required>
            </div>
            <div class="mb-6">
              <label class="block text-sm text-slate-400 mb-1">Mot de passe</label>
              <input type="password" id="password" class="w-full bg-slate-700 text-white rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="⬢⬢⬢⬢⬢⬢⬢⬢" required minlength="8">
            </div>
            <button type="submit" class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 rounded-lg transition">
              Se connecter
            </button>
          </form>
        </div>

        <p class="text-center text-slate-500 text-sm mt-4">
          �S�️ Apprenez où que vous soyez, même sans connexion
        </p>
      </div>
    </div>
  `;
}

function bindLoginEvents() {
  let isLogin = true;

  document.getElementById("tab-login")?.addEventListener("click", () => {
    isLogin = true;
    document.getElementById("tab-login").className = "flex-1 py-2 text-center font-medium rounded-l-lg bg-blue-600 text-white";
    document.getElementById("tab-register").className = "flex-1 py-2 text-center font-medium rounded-r-lg bg-slate-700 text-slate-300";
    document.querySelector("#auth-form button").textContent = "Se connecter";
  });

  document.getElementById("tab-register")?.addEventListener("click", () => {
    isLogin = false;
    document.getElementById("tab-register").className = "flex-1 py-2 text-center font-medium rounded-r-lg bg-blue-600 text-white";
    document.getElementById("tab-login").className = "flex-1 py-2 text-center font-medium rounded-l-lg bg-slate-700 text-slate-300";
    document.querySelector("#auth-form button").textContent = "Créer un compte";
  });

  document.getElementById("auth-form")?.addEventListener("submit", (e) => {
    e.preventDefault();
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    if (isLogin) {
      login(email, password);
    } else {
      register(email, password);
    }
  });
}

// ======================== NAVBAR ========================
function renderNav() {
  return `
    <nav class="bg-slate-800 border-b border-slate-700 sticky top-0 z-50">
      <div class="max-w-6xl mx-auto px-4">
        <div class="flex items-center justify-between h-14">
          <div class="flex items-center gap-2">
            <span class="text-xl">�S�️</span>
            <span class="font-bold text-white">AeroPath</span>
          </div>
          <div class="flex items-center gap-1">
            <button onclick="navigate('dashboard')" class="nav-btn px-3 py-1.5 rounded-lg text-sm ${state.view === 'dashboard' ? 'bg-blue-600 text-white' : 'text-slate-300 hover:bg-slate-700'}">Accueil</button>
            <button onclick="navigate('questions')" class="nav-btn px-3 py-1.5 rounded-lg text-sm ${state.view === 'questions' ? 'bg-blue-600 text-white' : 'text-slate-300 hover:bg-slate-700'}">Questions</button>
            <button onclick="navigate('lessons')" class="nav-btn px-3 py-1.5 rounded-lg text-sm ${state.view === 'lessons' ? 'bg-blue-600 text-white' : 'text-slate-300 hover:bg-slate-700'}">Leçons</button>
            <button onclick="navigate('history')" class="nav-btn px-3 py-1.5 rounded-lg text-sm ${state.view === 'history' ? 'bg-blue-600 text-white' : 'text-slate-300 hover:bg-slate-700'}">Historique</button>
            <button onclick="navigate('stats')" class="nav-btn px-3 py-1.5 rounded-lg text-sm ${state.view === 'stats' ? 'bg-blue-600 text-white' : 'text-slate-300 hover:bg-slate-700'}">Stats</button>
            <button onclick="navigate('recommendations')" class="nav-btn px-3 py-1.5 rounded-lg text-sm ${state.view === 'recommendations' ? 'bg-blue-600 text-white' : 'text-slate-300 hover:bg-slate-700'}">Recommandations</button>
            ${renderLanguageSelector()}
            <button onclick="logout()" class="ml-2 px-3 py-1.5 rounded-lg text-sm text-red-400 hover:bg-red-900/30">Déconnexion</button>
          </div>

        </div>
      </div>
    </nav>
  `;
}

// ======================== DASHBOARD ========================
function renderDashboard() {
  return `
    ${renderNav()}
    <div class="max-w-6xl mx-auto p-4">
      <div class="mb-6">
        <h1 class="text-2xl font-bold text-white">Bonjour ${currentUser?.email || "pilote"} �x9</h1>
        <p class="text-slate-400">Prêt à réviser ?</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
        <div class="bg-slate-800 rounded-xl p-4">
          <div class="text-3xl mb-2">�xa</div>
          <div class="text-2xl font-bold text-white" id="dash-questions">-</div>
          <div class="text-sm text-slate-400">Questions</div>
        </div>
        <div class="bg-slate-800 rounded-xl p-4">
          <div class="text-3xl mb-2">�x</div>
          <div class="text-2xl font-bold text-white" id="dash-lessons">-</div>
          <div class="text-sm text-slate-400">Leçons</div>
        </div>
        <div class="bg-slate-800 rounded-xl p-4">
          <div class="text-3xl mb-2">�S&</div>
          <div class="text-2xl font-bold text-white" id="dash-answers">-</div>
          <div class="text-sm text-slate-400">Réponses</div>
        </div>
      </div>

      <div class="bg-slate-800 rounded-xl p-4 mb-4">
        <h2 class="text-lg font-bold text-white mb-3">Actions rapides</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
          <button onclick="startRandomQuiz()" class="bg-blue-600 hover:bg-blue-700 text-white rounded-lg p-3 text-center transition">
            <div class="text-2xl mb-1">�x}�</div>
            <div class="text-sm">Question aléatoire</div>
          </button>
          <button onclick="navigate('questions')" class="bg-purple-600 hover:bg-purple-700 text-white rounded-lg p-3 text-center transition">
            <div class="text-2xl mb-1">�x9</div>
            <div class="text-sm">Toutes les questions</div>
          </button>
          <button onclick="navigate('lessons')" class="bg-green-600 hover:bg-green-700 text-white rounded-lg p-3 text-center transition">
            <div class="text-2xl mb-1">�x</div>
            <div class="text-sm">Leçons</div>
          </button>
          <button onclick="navigate('recommendations')" class="bg-amber-600 hover:bg-amber-700 text-white rounded-lg p-3 text-center transition">
            <div class="text-2xl mb-1">�x}�</div>
            <div class="text-sm">Recommandations</div>
          </button>
        </div>
      </div>

      <div id="dash-recommendations" class="bg-slate-800 rounded-xl p-4">
        <h2 class="text-lg font-bold text-white mb-3">Recommandations</h2>
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadDashboardData() {
  const [qCount, lCount, stats] = await Promise.all([
    api("/api/questions/count"),
    api("/api/lessons/count"),
    api("/api/admin/stats"),
  ]);

  if (qCount) document.getElementById("dash-questions").textContent = qCount.count ?? "-";
  if (lCount) document.getElementById("dash-lessons").textContent = lCount.count ?? "-";
  if (stats) document.getElementById("dash-answers").textContent = stats.answers ?? "-";

  const recs = await api("/api/recommendations");
  if (recs && !recs.offline) {
    state.recommendations = recs;
    const el = document.getElementById("dash-recommendations");
    if (el) {
      el.innerHTML = `
        <h2 class="text-lg font-bold text-white mb-3">Recommandations</h2>
        <div class="space-y-2">
          <div class="flex justify-between text-sm">
            <span class="text-slate-400">Progression</span>
            <span class="text-white font-medium">${Math.round(recs.Progression || 0)}%</span>
          </div>
          <div class="w-full bg-slate-700 rounded-full h-2">
            <div class="bg-blue-500 rounded-full h-2" style="width:${Math.round(recs.Progression || 0)}%"></div>
          </div>
          <div class="flex justify-between text-sm mt-3">
            <span class="text-slate-400">Sujets faibles</span>
            <span class="text-red-400 font-medium">${recs.WeakTopics?.length || 0}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-slate-400">Cartes à réviser</span>
            <span class="text-amber-400 font-medium">${recs.DueCards?.length || 0}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-slate-400">Prochain palier</span>
            <span class="text-green-400 font-medium">${recs.NextMilestone || "-"}</span>
          </div>
        </div>
      `;
    }
  }
}

// ======================== LICENCES & CAT�0GORIES ========================
const LICENSES = [
  { id: "PPL", label: "PPL", icon: "�x:�️", desc: "Pilote Privé" },
  { id: "LAPL", label: "LAPL", icon: "�x:�️", desc: "Pilote Privé Léger" },
  { id: "CPL", label: "CPL", icon: "�S�️", desc: "Pilote Professionnel" },
  { id: "ATPL", label: "ATPL", icon: "�x:�", desc: "Transport Aérien" },
  { id: "IR", label: "IR", icon: "�x:�", desc: "Vol aux Instruments" },
];

const CATEGORIES = [
  { id: "meteorology", label: "Météorologie", icon: "�xR�️" },
  { id: "navigation", label: "Navigation", icon: "�x��" },
  { id: "airlaw", label: "Réglementation", icon: "�a️" },
  { id: "aircraft_general", label: "Connaissance Aéronef", icon: "�x�" },
  { id: "performance", label: "Performance", icon: "�x�" },
  { id: "human_performance", label: "Facteurs Humains", icon: "�x��" },
  { id: "operational_procedures", label: "Procédures", icon: "�x9" },
  { id: "communications", label: "Communications", icon: "�x�" },
  { id: "principles_of_flight", label: "Principes du Vol", icon: "�x:�️" },
  { id: "flight_planning", label: "Planification", icon: "�x9" },
  { id: "instrumentation", label: "Instruments", icon: "�xx" },
  { id: "emergency", label: "Urgences", icon: "�x �" },
  { id: "mass_and_balance", label: "Masse & Centrage", icon: "�a️" },
  { id: "radio_procedure", label: "Procédures Radio", icon: "�x�" },
];

// ======================== QUESTIONS ========================
function renderQuestions() {
  return `
    ${renderNav()}
    <div class="max-w-6xl mx-auto p-4">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-bold text-white">Questions</h1>
        <button onclick="startRandomQuiz()" class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg text-sm">�x}� Question aléatoire</button>
      </div>

      <!-- Grille des licences -->
      <h2 class="text-lg font-semibold text-slate-300 mb-3">Choisis une licence</h2>
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-3 mb-8">
        ${LICENSES.map(l => `
          <button onclick="navigate('questions-license', {license: '${l.id}'}); setTimeout(() => loadQuestionsByLicense('${l.id}'), 50)" class="bg-slate-800 hover:bg-slate-700 rounded-xl p-4 text-center transition border border-slate-700 hover:border-blue-500/50">
            <div class="text-3xl mb-2">${l.icon}</div>
            <div class="text-white font-bold">${l.label}</div>
            <div class="text-xs text-slate-400">${l.desc}</div>
          </button>
        `).join("")}
      </div>

      <!-- Questions récentes -->
      <h2 class="text-lg font-semibold text-slate-300 mb-3">Dernières questions</h2>
      <div id="questions-list" class="space-y-3">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadQuestions() {
  const data = await api("/api/questions?limit=10");
  if (data?.questions) {
    state.questions = data.questions;
    renderQuestionsList(data.questions);
  }
}

function renderQuestionsList(questions) {
  const el = document.getElementById("questions-list");
  if (!el) return;

  if (!questions || questions.length === 0) {
    el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune question trouvée</p>';
    return;
  }

  el.innerHTML = questions.map((q) => `
    <div class="bg-slate-800 rounded-xl p-4 cursor-pointer hover:bg-slate-750 transition" onclick="showQuestionDetail('${q.id}')">
      <div class="flex justify-between items-start">
        <div class="flex-1">
          <p class="text-white font-medium">${q.question_fr || q.question_en}</p>
          <div class="flex gap-2 mt-2">
            <span class="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">${q.license || "-"}</span>
            <span class="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">${q.category || "-"}</span>
            <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${q.difficulty || "?"}</span>
            <span class="text-xs bg-green-900 text-green-300 px-2 py-0.5 rounded">${q.theme || "-"}</span>
          </div>
        </div>
        <span class="text-slate-500 ml-2">⬺</span>
      </div>
    </div>
  `).join("");
}

// ======================== QUESTIONS PAR LICENCE ========================
function renderQuestionsByLicense() {
  const licenseId = state.currentLicense;
  const license = LICENSES.find(l => l.id === licenseId) || { id: licenseId, label: licenseId, icon: "�x9", desc: "" };
  return `
    ${renderNav()}
    <div class="max-w-6xl mx-auto p-4">
      <button onclick="navigate('questions')" class="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
        � � Retour aux licences
      </button>
      <div class="flex items-center gap-3 mb-6">
        <span class="text-4xl">${license.icon}</span>
        <div>
          <h1 class="text-2xl font-bold text-white">${license.label}</h1>
          <p class="text-slate-400">${license.desc}</p>
        </div>
      </div>

      <h2 class="text-lg font-semibold text-slate-300 mb-3">Choisis une catégorie</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3 mb-8">
        ${CATEGORIES.map(c => `
          <button onclick="navigate('questions-category', {license: '${licenseId}', category: '${c.id}'}); setTimeout(() => loadQuestionsByCategory('${licenseId}', '${c.id}'), 50)" class="bg-slate-800 hover:bg-slate-700 rounded-xl p-4 text-left transition border border-slate-700 hover:border-purple-500/50">
            <div class="flex items-center gap-3">
              <span class="text-2xl">${c.icon}</span>
              <div>
                <div class="text-white font-medium">${c.label}</div>
                <div class="text-xs text-slate-500" id="q-count-${licenseId}-${c.id}">Chargement...</div>
              </div>
            </div>
          </button>
        `).join("")}
      </div>

      <h2 class="text-lg font-semibold text-slate-300 mb-3">Toutes les questions ${license.label}</h2>
      <div id="questions-license-list" class="space-y-3">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadQuestionsByLicense(licenseId) {
  state.currentLicense = licenseId;
  const data = await api(`/api/questions/by-license/${licenseId}`);
  if (data?.questions) {
    renderQuestionsLicenseList(data.questions);
  }
  // Compter les questions par catégorie depuis les données déjà récupérées
  if (data?.questions) {
    CATEGORIES.forEach((c) => {
      const count = data.questions.filter(q => q.category === c.id).length;
      const el = document.getElementById(`q-count-${licenseId}-${c.id}`);
      if (el) el.textContent = `${count} questions`;
    });
  }
}

function renderQuestionsLicenseList(questions) {
  const el = document.getElementById("questions-license-list");
  if (!el) return;
  if (!questions || questions.length === 0) {
    el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune question pour cette licence</p>';
    return;
  }
  el.innerHTML = questions.map((q) => `
    <div class="bg-slate-800 rounded-xl p-4 cursor-pointer hover:bg-slate-750 transition" onclick="showQuestionDetail('${q.id}')">
      <div class="flex justify-between items-start">
        <div class="flex-1">
          <p class="text-white font-medium">${q.question_fr || q.question_en}</p>
          <div class="flex gap-2 mt-2">
            <span class="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">${q.category || "-"}</span>
            <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${q.difficulty || "?"}</span>
            <span class="text-xs bg-green-900 text-green-300 px-2 py-0.5 rounded">${q.theme || "-"}</span>
          </div>
        </div>
        <span class="text-slate-500 ml-2">⬺</span>
      </div>
    </div>
  `).join("");
}

// ======================== QUESTIONS PAR CAT�0GORIE ========================
function renderQuestionsByCategory() {
  const licenseId = state.currentLicense;
  const catId = state.currentCategory;
  const license = LICENSES.find(l => l.id === licenseId) || { id: licenseId, label: licenseId, icon: "�x9" };
  const cat = CATEGORIES.find(c => c.id === catId) || { id: catId, label: catId, icon: "�x9" };
  return `
    ${renderNav()}
    <div class="max-w-6xl mx-auto p-4">
      <button onclick="navigate('questions-license'); setTimeout(() => loadQuestionsByLicense('${licenseId}'), 50)" class="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
        � � Retour à ${license.label}
      </button>
      <div class="flex items-center gap-3 mb-6">
        <span class="text-4xl">${cat.icon}</span>
        <div>
          <h1 class="text-2xl font-bold text-white">${cat.label}</h1>
          <p class="text-slate-400">${license.icon} ${license.label}</p>
        </div>
      </div>
      <div class="flex gap-2 mb-4">
        <button onclick="startCategoryQuiz('${licenseId}', '${catId}')" class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg text-sm">�x}� Quiz sur cette catégorie</button>
      </div>
      <div id="questions-category-list" class="space-y-3">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadQuestionsByCategory(licenseId, categoryId) {
  state.currentLicense = licenseId;
  state.currentCategory = categoryId;
  const data = await api(`/api/questions/by-license/${licenseId}/category/${categoryId}`);
  if (data?.questions) {
    renderQuestionsCategoryList(data.questions);
  }
}

function renderQuestionsCategoryList(questions) {
  const el = document.getElementById("questions-category-list");
  if (!el) return;
  if (!questions || questions.length === 0) {
    el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune question dans cette catégorie</p>';
    return;
  }
  el.innerHTML = questions.map((q) => `
    <div class="bg-slate-800 rounded-xl p-4 cursor-pointer hover:bg-slate-750 transition" onclick="showQuestionDetail('${q.id}')">
      <div class="flex justify-between items-start">
        <div class="flex-1">
          <p class="text-white font-medium">${q.question_fr || q.question_en}</p>
          <div class="flex gap-2 mt-2">
            <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${q.difficulty || "?"}</span>
            <span class="text-xs bg-green-900 text-green-300 px-2 py-0.5 rounded">${q.theme || "-"}</span>
          </div>
        </div>
        <span class="text-slate-500 ml-2">⬺</span>
      </div>
    </div>
  `).join("");
}

async function startCategoryQuiz(licenseId, categoryId) {
  const data = await api(`/api/exam/license/${licenseId}/category/${categoryId}`);
  if (data?.questions) {
    state.quizQuestions = data.questions;
    state.quizIndex = 0;
    state.quizScore = 0;
    state.quizFeedback = null;
    navigate("quiz");
  }
}

async function showQuestionDetail(id) {
  state.currentQuestion = id;
  navigate("question-detail", {
    license: state.currentLicense,
    category: state.currentCategory,
    questionId: id,
  });
}

function renderQuestionDetail() {
  const id = state.currentQuestion;
  return `
    ${renderNav()}
    <div class="max-w-3xl mx-auto p-4">
      <button onclick="goBack()" class="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
        � � Retour
      </button>
      <div id="question-detail-content" class="bg-slate-800 rounded-xl p-6">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

// Charger le détail de la question après le rendu
document.addEventListener("DOMContentLoaded", () => {
  // On utilise un MutationObserver pour détecter quand la page question-detail est rendue
  const observer = new MutationObserver(() => {
    if (state.view === "question-detail" && state.currentQuestion) {
      loadQuestionDetail(state.currentQuestion);
    }
  });
  observer.observe(document.getElementById("app"), { childList: true, subtree: true });
});

async function loadQuestionDetail(id) {
  const data = await api(`/api/questions/${id}`);
  const el = document.getElementById("question-detail-content");
  if (!el) return;
  if (!data) {
    el.innerHTML = '<p class="text-red-400">Erreur de chargement</p>';
    return;
  }

  el.innerHTML = `
    <div class="flex justify-between items-start mb-4">
      <h3 class="text-lg font-bold text-white">Question</h3>
    </div>
    <p class="text-white text-lg mb-6">${data.question_fr || data.question_en}</p>
    <div class="space-y-3 mb-6">
      ${(data.options || []).map((opt, i) => `
        <button onclick="selectOption('${id}', '${String.fromCharCode(65 + i)}', this)" class="w-full bg-slate-700 hover:bg-slate-600 text-white text-left rounded-lg p-3 transition border border-transparent hover:border-blue-500/50" data-option="${String.fromCharCode(65 + i)}">
          ${String.fromCharCode(65 + i)}. ${opt}
        </button>
      `).join("")}
    </div>
    <div class="flex gap-2 flex-wrap mb-6">
      <span class="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">${data.license || "-"}</span>
      <span class="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">${data.category || "-"}</span>
      <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${data.difficulty || "?"}</span>
      <span class="text-xs bg-green-900 text-green-300 px-2 py-0.5 rounded">${data.theme || "-"}</span>
    </div>
    <div id="answer-result-${id}" class="mt-4 hidden"></div>
  `;
}

async function selectOption(questionId, selected, btn) {
  // Désactiver tous les boutons
  document.querySelectorAll(`#question-detail-content button[data-option]`).forEach(b => {
    b.disabled = true;
    b.className = b.className.replace('hover:bg-slate-600', '');
  });

  // Marquer la sélection
  btn.className = btn.className.replace('border-transparent', 'border-blue-500');

  // Utiliser l'API de vérification de réponse (POST /api/questions/answer)
  const data = await api("/api/questions/answer", {
    method: "POST",
    body: JSON.stringify({ question_id: questionId, answer: selected }),
  });
  if (!data) return;

  const resultEl = document.getElementById(`answer-result-${questionId}`);
  if (!resultEl) return;

  const isCorrect = data.correct;

  // Colorer les boutons
  document.querySelectorAll(`#question-detail-content button[data-option]`).forEach(b => {
    const opt = b.getAttribute('data-option');
    if (opt === data.correct_answer) {
      b.className = b.className.replace('bg-slate-700', 'bg-green-700');
      b.className = b.className.replace('border-transparent', 'border-green-500');
    } else if (opt === selected && !isCorrect) {
      b.className = b.className.replace('bg-slate-700', 'bg-red-700');
      b.className = b.className.replace('border-transparent', 'border-red-500');
    }
  });

  resultEl.className = `mt-4 p-4 rounded-lg ${isCorrect ? 'bg-green-900/30' : 'bg-red-900/30'}`;
  resultEl.innerHTML = `
    <p class="${isCorrect ? 'text-green-400' : 'text-red-400'} font-bold text-lg mb-2">
      ${isCorrect ? '�S& Correct !' : '�R Faux'}
    </p>
    <p class="text-green-400 font-medium mb-2">Réponse correcte : ${data.correct_answer || "?"}</p>
    ${data.explanation_fr ? `<p class="text-slate-300 mt-2">${data.explanation_fr}</p>` : ""}
    ${data.explanation_en ? `<p class="text-slate-300 mt-2">${data.explanation_en}</p>` : ""}
  `;
}

// ======================== QUIZ AL�0ATOIRE ========================
async function startRandomQuiz() {
  const data = await api("/api/questions/random?limit=5");
  if (data?.questions) {
    state.quizQuestions = data.questions;
    state.quizIndex = 0;
    state.quizScore = 0;
    state.quizFeedback = null;
    navigate("quiz");
  }
}

function renderQuiz() {
  if (state.quizIndex >= state.quizQuestions.length) {
    state.quizFeedback = null;
    return renderQuizResult();
  }

  const q = state.quizQuestions[state.quizIndex];
  const fb = state.quizFeedback;

  return `
    ${renderNav()}
    <div class="max-w-2xl mx-auto p-4">
      <div class="mb-4">
        <div class="flex justify-between text-sm text-slate-400 mb-2">
          <span>Question ${state.quizIndex + 1}/${state.quizQuestions.length}</span>
          <span>Score: ${state.quizScore}/${state.quizIndex}</span>
        </div>
        <div class="w-full bg-slate-700 rounded-full h-2">
          <div class="bg-blue-500 rounded-full h-2" style="width:${(state.quizIndex / state.quizQuestions.length) * 100}%"></div>
        </div>
      </div>

      <div class="bg-slate-800 rounded-xl p-6">
        <p class="text-white text-lg mb-6">${q.question_fr || q.question_en}</p>
        <div class="space-y-3 mb-6">
          ${(q.options || []).map((opt, i) => {
            const letter = String.fromCharCode(65 + i);
            const isCorrectOpt = fb && letter === fb.correct_answer;
            const isSelected = fb && fb.selected_answer === letter;
            const isWrong = fb && isSelected && !fb.correct;
            let btnClass = fb ? 'bg-slate-700 text-white text-left rounded-lg p-3 cursor-default' : 'bg-slate-700 hover:bg-slate-600 text-white text-left rounded-lg p-3 transition';
            if (isCorrectOpt) btnClass = 'bg-green-700 text-white text-left rounded-lg p-3 border border-green-500';
            else if (isWrong) btnClass = 'bg-red-700 text-white text-left rounded-lg p-3 border border-red-500';
            return `
              <button onclick="${fb ? '' : `submitQuizAnswer('${q.id}', '${letter}')`}" class="${btnClass}" ${fb ? 'disabled' : ''}>
                ${letter}. ${opt}
              </button>
            `;
          }).join("")}
        </div>

        ${fb ? `
          <div class="p-4 rounded-lg ${fb.correct ? 'bg-green-900/30' : 'bg-red-900/30'} mb-4">
            <p class="${fb.correct ? 'text-green-400' : 'text-red-400'} font-bold text-lg mb-2">
              ${fb.correct ? '�S& Correct !' : '�R Faux'}
            </p>
            <p class="text-green-400 font-medium mb-2">Réponse correcte : ${fb.correct_answer}</p>
            ${fb.explanation_fr ? `<p class="text-slate-300 mt-2">${fb.explanation_fr}</p>` : ""}
            ${fb.explanation_en ? `<p class="text-slate-300 mt-2">${fb.explanation_en}</p>` : ""}
          </div>
          <button onclick="nextQuizQuestion()" class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 rounded-lg transition text-lg">
            ${state.quizIndex + 1 >= state.quizQuestions.length ? '�x` Voir les résultats' : '�~�️ Question suivante'}
          </button>
        ` : ''}
      </div>
    </div>
  `;
}

async function submitQuizAnswer(questionId, answer) {
  const data = await api("/api/questions/answer", {
    method: "POST",
    body: JSON.stringify({ question_id: questionId, answer }),
  });

  if (data?.correct) state.quizScore++;
  state.quizFeedback = {
    correct: data?.correct ?? false,
    correct_answer: data?.correct_answer ?? "?",
    selected_answer: answer,
    explanation_fr: data?.explanation_fr ?? "",
    explanation_en: data?.explanation_en ?? "",
  };
  render();
}

function nextQuizQuestion() {
  state.quizFeedback = null;
  state.quizIndex++;
  render();
}

function renderQuizResult() {
  const pct = state.quizQuestions.length > 0
    ? Math.round((state.quizScore / state.quizQuestions.length) * 100)
    : 0;

  return `
    ${renderNav()}
    <div class="max-w-2xl mx-auto p-4 text-center">
      <div class="bg-slate-800 rounded-xl p-8">
        <div class="text-6xl mb-4">${pct >= 80 ? "�x}0" : pct >= 50 ? "�x�" : "�xa"}</div>
        <h2 class="text-2xl font-bold text-white mb-2">Quiz terminé !</h2>
        <p class="text-4xl font-bold text-blue-400 mb-2">${pct}%</p>
        <p class="text-slate-400 mb-6">${state.quizScore}/${state.quizQuestions.length} bonnes réponses</p>
        <div class="flex gap-3 justify-center">
          <button onclick="startRandomQuiz()" class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg">Rejouer</button>
          <button onclick="navigate('dashboard')" class="bg-slate-700 hover:bg-slate-600 text-white px-6 py-2 rounded-lg">Accueil</button>
        </div>
      </div>
    </div>
  `;
}

// ======================== HISTORIQUE ========================
function renderHistory() {
  return `
    ${renderNav()}
    <div class="max-w-4xl mx-auto p-4">
      <h1 class="text-2xl font-bold text-white mb-4">Historique</h1>
      <div id="history-list" class="space-y-2">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadHistory() {
  const data = await api("/api/history?limit=100");
  if (data?.entries) {
    state.history = data.entries;
    const el = document.getElementById("history-list");
    if (!el) return;

    if (data.entries.length === 0) {
      el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucun historique</p>';
      return;
    }

    el.innerHTML = data.entries.map((h) => `
      <div class="bg-slate-800 rounded-lg p-3 flex justify-between items-center">
        <div class="flex-1 min-w-0">
          <p class="text-white text-sm truncate">${h.question_fr || h.question_en}</p>
          <div class="flex gap-2 mt-1">
            <span class="text-xs text-slate-500">${h.theme || "-"}</span>
            <span class="text-xs text-slate-500">Niv. ${h.difficulty || "?"}</span>
          </div>
        </div>
        <span class="ml-3 ${h.was_correct ? 'text-green-400' : 'text-red-400'} font-medium text-sm">
          ${h.was_correct ? "�S&" : "�R"}
        </span>
      </div>
    `).join("");
  }
}

// ======================== LE�!ONS ========================
function renderLessons() {
  return `
    ${renderNav()}
    <div class="max-w-6xl mx-auto p-4">
      <h1 class="text-2xl font-bold text-white mb-6">Leçons</h1>

      <!-- Grille des licences -->
      <h2 class="text-lg font-semibold text-slate-300 mb-3">Choisis une licence</h2>
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-3 mb-8">
        ${LICENSES.map(l => `
          <button onclick="navigate('lessons-license', {license: '${l.id}'}); setTimeout(() => loadLessonsByLicense('${l.id}'), 50)" class="bg-slate-800 hover:bg-slate-700 rounded-xl p-4 text-center transition border border-slate-700 hover:border-green-500/50">
            <div class="text-3xl mb-2">${l.icon}</div>
            <div class="text-white font-bold">${l.label}</div>
            <div class="text-xs text-slate-400">${l.desc}</div>
          </button>
        `).join("")}
      </div>

      <!-- Dernières leçons -->
      <h2 class="text-lg font-semibold text-slate-300 mb-3">Dernières leçons</h2>
      <div id="lessons-list" class="space-y-3">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadLessons() {
  const data = await api("/api/lessons?limit=10");
  if (data?.data) {
    state.lessons = data.data;
    renderLessonsList(data.data);
  } else if (Array.isArray(data)) {
    state.lessons = data;
    renderLessonsList(data);
  } else {
    const el = document.getElementById("lessons-list");
    if (el) el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune leçon trouvée</p>';
  }
}

function renderLessonsList(lessons) {
  const el = document.getElementById("lessons-list");
  if (!el) return;

  if (!lessons || lessons.length === 0) {
    el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune leçon trouvée</p>';
    return;
  }

  el.innerHTML = lessons.map((l) => `
    <div class="bg-slate-800 rounded-xl p-4 cursor-pointer hover:bg-slate-750 transition" onclick="showLessonDetail('${l.id}')">
      <h3 class="text-white font-medium">${l.title_fr || l.title_en}</h3>
      <div class="flex gap-2 mt-2">
        <span class="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">${l.license || "-"}</span>
        <span class="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">${l.category || "-"}</span>
        <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${l.difficulty || "?"}</span>
      </div>
    </div>
  `).join("");
}

// ======================== LE�!ONS PAR LICENCE ========================
function renderLessonsByLicense() {
  const licenseId = state.currentLicense;
  const license = LICENSES.find(l => l.id === licenseId) || { id: licenseId, label: licenseId, icon: "�x9", desc: "" };
  return `
    ${renderNav()}
    <div class="max-w-6xl mx-auto p-4">
      <button onclick="navigate('lessons')" class="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
        � � Retour aux licences
      </button>
      <div class="flex items-center gap-3 mb-6">
        <span class="text-4xl">${license.icon}</span>
        <div>
          <h1 class="text-2xl font-bold text-white">${license.label}</h1>
          <p class="text-slate-400">${license.desc}</p>
        </div>
      </div>

      <h2 class="text-lg font-semibold text-slate-300 mb-3">Choisis une catégorie</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3 mb-8">
        ${CATEGORIES.map(c => `
          <button onclick="navigate('lessons-category', {license: '${licenseId}', category: '${c.id}'}); setTimeout(() => loadLessonsByCategory('${licenseId}', '${c.id}'), 50)" class="bg-slate-800 hover:bg-slate-700 rounded-xl p-4 text-left transition border border-slate-700 hover:border-green-500/50">
            <div class="flex items-center gap-3">
              <span class="text-2xl">${c.icon}</span>
              <div>
                <div class="text-white font-medium">${c.label}</div>
                <div class="text-xs text-slate-500" id="l-count-${licenseId}-${c.id}">Chargement...</div>
              </div>
            </div>
          </button>
        `).join("")}
      </div>

      <h2 class="text-lg font-semibold text-slate-300 mb-3">Toutes les leçons ${license.label}</h2>
      <div id="lessons-license-list" class="space-y-3">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadLessonsByLicense(licenseId) {
  state.currentLicense = licenseId;
  const data = await api(`/api/lessons/license/${licenseId}`);
  if (Array.isArray(data)) {
    renderLessonsLicenseList(data);
  } else if (data?.data) {
    renderLessonsLicenseList(data.data);
  } else {
    const el = document.getElementById("lessons-license-list");
    if (el) el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune leçon pour cette licence</p>';
  }
  // Compter les leçons par catégorie depuis les données déjà récupérées
  const lessonsArr = Array.isArray(data) ? data : (data?.data || []);
  if (lessonsArr.length > 0) {
    CATEGORIES.forEach((c) => {
      const count = lessonsArr.filter(l => l.category === c.id).length;
      const el = document.getElementById(`l-count-${licenseId}-${c.id}`);
      if (el) el.textContent = `${count} leçons`;
    });
  }
}

function renderLessonsLicenseList(lessons) {
  const el = document.getElementById("lessons-license-list");
  if (!el) return;
  if (!lessons || lessons.length === 0) {
    el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune leçon pour cette licence</p>';
    return;
  }
  el.innerHTML = lessons.map((l) => `
    <div class="bg-slate-800 rounded-xl p-4 cursor-pointer hover:bg-slate-750 transition" onclick="showLessonDetail('${l.id}')">
      <h3 class="text-white font-medium">${l.title_fr || l.title_en}</h3>
      <div class="flex gap-2 mt-2">
        <span class="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">${l.category || "-"}</span>
        <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${l.difficulty || "?"}</span>
      </div>
    </div>
  `).join("");
}

// ======================== LE�!ONS PAR CAT�0GORIE ========================
function renderLessonsByCategory() {
  const licenseId = state.currentLicense;
  const catId = state.currentCategory;
  const license = LICENSES.find(l => l.id === licenseId) || { id: licenseId, label: licenseId, icon: "�x9" };
  const cat = CATEGORIES.find(c => c.id === catId) || { id: catId, label: catId, icon: "�x9" };
  return `
    ${renderNav()}
    <div class="max-w-6xl mx-auto p-4">
      <button onclick="navigate('lessons-license'); setTimeout(() => loadLessonsByLicense('${licenseId}'), 50)" class="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
        � � Retour à ${license.label}
      </button>
      <div class="flex items-center gap-3 mb-6">
        <span class="text-4xl">${cat.icon}</span>
        <div>
          <h1 class="text-2xl font-bold text-white">${cat.label}</h1>
          <p class="text-slate-400">${license.icon} ${license.label}</p>
        </div>
      </div>
      <div id="lessons-category-list" class="space-y-3">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadLessonsByCategory(licenseId, categoryId) {
  state.currentLicense = licenseId;
  state.currentCategory = categoryId;
  const data = await api(`/api/lessons/by-license/${licenseId}/category/${categoryId}`);
  if (Array.isArray(data)) {
    renderLessonsCategoryList(data);
  } else if (data?.data) {
    renderLessonsCategoryList(data.data);
  } else {
    const el = document.getElementById("lessons-category-list");
    if (el) el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune leçon dans cette catégorie</p>';
  }
}

function renderLessonsCategoryList(lessons) {
  const el = document.getElementById("lessons-category-list");
  if (!el) return;
  if (!lessons || lessons.length === 0) {
    el.innerHTML = '<p class="text-slate-400 text-center py-8">Aucune leçon dans cette catégorie</p>';
    return;
  }
  el.innerHTML = lessons.map((l) => `
    <div class="bg-slate-800 rounded-xl p-4 cursor-pointer hover:bg-slate-750 transition" onclick="showLessonDetail('${l.id}')">
      <h3 class="text-white font-medium">${l.title_fr || l.title_en}</h3>
      <div class="flex gap-2 mt-2">
        <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${l.difficulty || "?"}</span>
      </div>
    </div>
  `).join("");
}

async function showLessonDetail(id) {
  state.currentLesson = id;
  navigate("lesson-detail", {
    license: state.currentLicense,
    category: state.currentCategory,
    lessonId: id,
  });
}

function renderLessonDetail() {
  const id = state.currentLesson;
  return `
    ${renderNav()}
    <div class="max-w-4xl mx-auto p-4">
      <button onclick="goBack()" class="text-slate-400 hover:text-white mb-4 flex items-center gap-1">
        � � Retour
      </button>
      <div id="lesson-detail-content" class="bg-slate-800 rounded-xl p-6">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

function renderMarkdown(text) {
  if (!text) return "";
  // �0chapper le HTML
  let html = text
    .replace(/&/g, "&")
    .replace(/</g, "<")
    .replace(/>/g, ">");
  // Titres
  html = html.replace(/^### (.+)$/gm, '<h3 class="text-lg font-bold text-white mt-4 mb-2">$1</h3>');
  html = html.replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold text-white mt-5 mb-2">$1</h2>');
  html = html.replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold text-white mt-6 mb-3">$1</h1>');
  // Gras et italique
  html = html.replace(/\*\*(.+?)\*\*/g, '<strong class="text-white font-bold">$1</strong>');
  html = html.replace(/\*(.+?)\*/g, '<em class="text-slate-200 italic">$1</em>');
  // Listes
  html = html.replace(/^- (.+)$/gm, '<li class="text-slate-300 ml-4 list-disc">$1</li>');
  html = html.replace(/^\d+\. (.+)$/gm, '<li class="text-slate-300 ml-4 list-decimal">$1</li>');
  // Paragraphes
  html = html.replace(/\n\n/g, '</p><p class="text-slate-300 mb-3 leading-relaxed">');
  html = '<p class="text-slate-300 mb-3 leading-relaxed">' + html + '</p>';
  return html;
}

async function loadLessonDetail(id) {
  const data = await api(`/api/lessons/${id}`);
  const el = document.getElementById("lesson-detail-content");
  if (!el) return;
  if (!data) {
    el.innerHTML = '<p class="text-red-400">Erreur de chargement</p>';
    return;
  }

  const content = data.content_fr || data.content_en || "";
  const renderedContent = content ? renderMarkdown(content) : '<p class="text-slate-400 italic">Contenu non disponible</p>';

  el.innerHTML = `
    <div class="flex justify-between items-start mb-4">
      <h3 class="text-xl font-bold text-white">${data.title_fr || data.title_en}</h3>
    </div>
    <div class="mb-6">${renderedContent}</div>
    <div class="flex gap-2 flex-wrap mb-6">
      <span class="text-xs bg-blue-900 text-blue-300 px-2 py-0.5 rounded">${data.license || "-"}</span>
      <span class="text-xs bg-purple-900 text-purple-300 px-2 py-0.5 rounded">${data.category || "-"}</span>
      <span class="text-xs bg-amber-900 text-amber-300 px-2 py-0.5 rounded">Niv. ${data.difficulty || "?"}</span>
    </div>
    <button onclick="startLessonQuiz('${id}')" class="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-3 rounded-lg transition text-lg">
      �x� Quiz sur cette leçon
    </button>
  `;
}

async function startLessonQuiz(lessonId) {
  const data = await api(`/api/lessons/${lessonId}/quiz`);
  if (data?.questions) {
    state.quizQuestions = data.questions;
    state.quizIndex = 0;
    state.quizScore = 0;
    state.quizFeedback = null;
    navigate("quiz");
  }
}

// ======================== STATISTIQUES ========================
function renderStats() {
  return `
    ${renderNav()}
    <div class="max-w-4xl mx-auto p-4">
      <h1 class="text-2xl font-bold text-white mb-4">Statistiques</h1>
      <div id="stats-content" class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadStats() {
  const [stats, adminStats] = await Promise.all([
    api("/api/stats"),
    api("/api/admin/stats"),
  ]);

  const el = document.getElementById("stats-content");
  if (!el) return;

  el.innerHTML = `
    <div class="bg-slate-800 rounded-xl p-4">
      <h3 class="text-white font-medium mb-3">�x` Mes statistiques</h3>
      ${stats ? `
        <div class="space-y-2 text-sm">
          ${Object.entries(stats).map(([k, v]) => `
            <div class="flex justify-between">
              <span class="text-slate-400">${k}</span>
              <span class="text-white">${v}</span>
            </div>
          `).join("")}
        </div>
      ` : '<p class="text-slate-400">Non disponible</p>'}
    </div>
    <div class="bg-slate-800 rounded-xl p-4">
      <h3 class="text-white font-medium mb-3">�xR� Global</h3>
      ${adminStats ? `
        <div class="space-y-2 text-sm">
          <div class="flex justify-between">
            <span class="text-slate-400">�0tudiants</span>
            <span class="text-white">${adminStats.students || 0}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-slate-400">Questions</span>
            <span class="text-white">${adminStats.questions || 0}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-slate-400">Réponses</span>
            <span class="text-white">${adminStats.answers || 0}</span>
          </div>
        </div>
      ` : '<p class="text-slate-400">Non disponible</p>'}
    </div>
  `;
}

// ======================== RECOMMANDATIONS ========================
function renderRecommendations() {
  return `
    ${renderNav()}
    <div class="max-w-4xl mx-auto p-4">
      <h1 class="text-2xl font-bold text-white mb-4">Recommandations</h1>
      <div id="recs-content" class="space-y-4">
        <p class="text-slate-400">Chargement...</p>
      </div>
    </div>
  `;
}

async function loadRecommendations() {
  const data = await api("/api/recommendations");
  const el = document.getElementById("recs-content");
  if (!el) return;

  if (!data || data.offline) {
    el.innerHTML = '<p class="text-slate-400">Non disponible hors-ligne</p>';
    return;
  }

  state.recommendations = data;

  el.innerHTML = `
    <div class="bg-slate-800 rounded-xl p-4">
      <h3 class="text-white font-medium mb-3">�x� Progression</h3>
      <div class="flex justify-between text-sm mb-1">
        <span class="text-slate-400">Globale</span>
        <span class="text-white font-medium">${Math.round(data.Progression || 0)}%</span>
      </div>
      <div class="w-full bg-slate-700 rounded-full h-3">
        <div class="bg-blue-500 rounded-full h-3 transition-all" style="width:${Math.round(data.Progression || 0)}%"></div>
      </div>
      <p class="text-sm text-slate-500 mt-2">Prochain palier : ${data.NextMilestone || "-"}</p>
    </div>

    <div class="bg-slate-800 rounded-xl p-4">
      <h3 class="text-white font-medium mb-3">�x}� Sujets à travailler</h3>
      ${(data.WeakTopics || []).length > 0 ? `
        <div class="space-y-2">
          ${data.WeakTopics.map((t) => `
            <div class="flex justify-between items-center">
              <span class="text-slate-300">${t.Theme || t.theme || t}</span>
              <span class="text-red-400 text-sm">${Math.round(t.Score || t.score || 0)}%</span>
            </div>
          `).join("")}
        </div>
      ` : '<p class="text-slate-400">Aucun sujet faible détecté</p>'}
    </div>

    <div class="bg-slate-800 rounded-xl p-4">
      <h3 class="text-white font-medium mb-3">�x& Cartes à réviser</h3>
      ${(data.DueCards || []).length > 0 ? `
        <div class="space-y-2">
          ${data.DueCards.slice(0, 10).map((c) => `
            <div class="bg-slate-700 rounded-lg p-3">
              <p class="text-white text-sm">${c.Question || c.question_fr || c.question_en || "Question"}</p>
              <p class="text-xs text-slate-500 mt-1">Prochaine révision : ${c.NextReview || c.next_review || "bientôt"}</p>
            </div>
          `).join("")}
          ${data.DueCards.length > 10 ? `<p class="text-sm text-slate-500 mt-2">Et ${data.DueCards.length - 10} autres...</p>` : ""}
        </div>
      ` : '<p class="text-slate-400">Tout est à jour ! �x}0</p>'}
    </div>

    <div class="bg-slate-800 rounded-xl p-4">
      <h3 class="text-white font-medium mb-3">�x�  Maîtrise par licence</h3>
      ${(data.MasteryByLicense || []).length > 0 ? `
        <div class="space-y-2">
          ${data.MasteryByLicense.map((m) => `
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-slate-300">${m.License || m.license}</span>
                <span class="text-white">${Math.round(m.Score || m.score || 0)}%</span>
              </div>
              <div class="w-full bg-slate-700 rounded-full h-2">
                <div class="bg-green-500 rounded-full h-2" style="width:${Math.round(m.Score || m.score || 0)}%"></div>
              </div>
            </div>
          `).join("")}
        </div>
      ` : '<p class="text-slate-400">Pas encore de données</p>'}
    </div>
  `;
}

// ======================== TOAST ========================
function showToast(message, type = "info") {
  const colors = {
    info: "bg-blue-600",
    error: "bg-red-600",
    success: "bg-green-600",
  };

  const toast = document.createElement("div");
  toast.className = `fixed bottom-4 right-4 ${colors[type] || colors.info} text-white px-4 py-2 rounded-lg shadow-lg z-50 transition-all duration-300`;
  toast.textContent = message;
  document.body.appendChild(toast);

  setTimeout(() => {
    toast.style.opacity = "0";
    setTimeout(() => toast.remove(), 300);
  }, 3000);
}

// ======================== INITIALISATION ========================
window.addEventListener("popstate", onPopState);

async function init() {
  var route = parseURL(window.location.pathname);
  var publicViews = ["home", "login", "login-form"];

  if (!authToken && !publicViews.includes(route.view)) {
    history.replaceState(null, "", "/");
    state.view = "home";
    render();
    return;
  }

  if (authToken) {
    await loadUser();
  }

  applyParsedRoute(route);

  if (authToken && route.view === "home") {
    navigate("dashboard");
    return;
  }

  render();
}

init();
