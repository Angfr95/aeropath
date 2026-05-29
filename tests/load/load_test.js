// Test de charge AeroPath avec k6
// Installation : https://k6.io/docs/getting-started/installation/
// Exécution : k6 run tests/load/load_test.js

import http from 'k6/http';
import { check, sleep, group } from 'k6';
import { Rate, Trend, Counter } from 'k6/metrics';

// ======================== MÉTRIQUES PERSONNALISÉES ========================

const errorRate = new Rate('errors');
const authDuration = new Trend('auth_duration');
const questionsDuration = new Trend('questions_duration');
const lessonsDuration = new Trend('lessons_duration');
const quizDuration = new Trend('quiz_duration');
const answerDuration = new Trend('answer_duration');
const recommendationsDuration = new Trend('recommendations_duration');
const adminDuration = new Trend('admin_duration');
const totalRequests = new Counter('total_requests');

// ======================== CONFIGURATION ========================

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';
const STAGES = [
  { duration: '30s', target: 10 },   // Montée à 10 utilisateurs
  { duration: '1m', target: 50 },     // Montée à 50 utilisateurs
  { duration: '2m', target: 100 },    // Montée à 100 utilisateurs
  { duration: '2m', target: 100 },    // Maintien à 100 utilisateurs
  { duration: '1m', target: 200 },    // Montée à 200 utilisateurs
  { duration: '2m', target: 200 },    // Maintien à 200 utilisateurs
  { duration: '1m', target: 0 },      // Descente
];

export const options = {
  stages: STAGES,
  thresholds: {
    http_req_duration: ['p(95)<2000'],  // 95% des requêtes < 2s
    http_req_failed: ['rate<0.05'],     // < 5% d'erreurs
    errors: ['rate<0.05'],              // < 5% d'erreurs métier
    auth_duration: ['p(95)<3000'],      // Auth < 3s
    questions_duration: ['p(95)<1500'], // Questions < 1.5s
    lessons_duration: ['p(95)<1500'],   // Leçons < 1.5s
    quiz_duration: ['p(95)<2000'],      // Quiz < 2s
    answer_duration: ['p(95)<2000'],    // Réponse < 2s
    recommendations_duration: ['p(95)<2000'], // Recommandations < 2s
  },
};

// ======================== ÉTAT PAR VU ========================

const vuState = {};

function getOrInitState(vu) {
  if (!vuState[vu]) {
    vuState[vu] = {
      token: null,
      studentId: null,
      questionIds: [],
      lessonIds: [],
    };
  }
  return vuState[vu];
}

// ======================== FONCTIONS UTILITAIRES ========================

function randomItem(arr) {
  return arr[Math.floor(Math.random() * arr.length)];
}

function checkResponse(name, response, expectedStatus = 200) {
  const success = check(response, {
    [`${name} status is ${expectedStatus}`]: (r) => r.status === expectedStatus,
    [`${name} body not empty`]: (r) => r.body.length > 0,
  });
  if (!success) {
    errorRate.add(1);
    console.error(`❌ ${name}: status=${response.status}, body=${response.body.substring(0, 200)}`);
  }
  totalRequests.add(1);
  return success;
}

// ======================== SCÉNARIOS ========================

function scenarioAuth(state) {
  group('Auth - Register', () => {
    const email = `loadtest-${__VU}-${Date.now()}@test.com`;
    const payload = JSON.stringify({
      email: email,
      password: 'password123',
      lang: __VU % 2 === 0 ? 'fr' : 'en',
    });

    const res = http.post(`${BASE_URL}/api/auth/register`, payload, {
      headers: { 'Content-Type': 'application/json' },
    });
    authDuration.add(res.timings.duration);

    if (checkResponse('Register', res, 201)) {
      const body = JSON.parse(res.body);
      state.token = body.token;
      state.studentId = body.student_id;
    }
  });

  if (!state.token) {
    // Fallback: login
    group('Auth - Login', () => {
      const payload = JSON.stringify({
        email: `loadtest-${__VU}@test.com`,
        password: 'password123',
      });

      const res = http.post(`${BASE_URL}/api/auth/login`, payload, {
        headers: { 'Content-Type': 'application/json' },
      });
      authDuration.add(res.timings.duration);

      if (checkResponse('Login', res)) {
        const body = JSON.parse(res.body);
        state.token = body.token;
        state.studentId = body.student_id;
      }
    });
  }
}

function scenarioQuestions(state) {
  if (!state.token) return;

  group('Questions - List', () => {
    const res = http.get(`${BASE_URL}/api/questions?page=1&page_size=20`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    questionsDuration.add(res.timings.duration);

    if (checkResponse('List Questions', res)) {
      try {
        const body = JSON.parse(res.body);
        if (body.data && body.data.length > 0) {
          state.questionIds = body.data.map(q => q.id);
        }
      } catch (e) {
        console.error('Parse error:', e);
      }
    }
  });

  if (state.questionIds.length > 0) {
    group('Questions - Get One', () => {
      const qId = randomItem(state.questionIds);
      const res = http.get(`${BASE_URL}/api/questions/${qId}`, {
        headers: { 'Authorization': `Bearer ${state.token}` },
      });
      questionsDuration.add(res.timings.duration);
      checkResponse('Get Question', res);
    });

    group('Questions - Answer', () => {
      const qId = randomItem(state.questionIds);
      const payload = JSON.stringify({
        question_id: qId,
        answer: String.fromCharCode(65 + Math.floor(Math.random() * 4)), // A, B, C, D
      });

      const res = http.post(`${BASE_URL}/api/questions/answer`, payload, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${state.token}`,
        },
      });
      answerDuration.add(res.timings.duration);
      checkResponse('Answer Question', res);
    });
  }
}

function scenarioLessons(state) {
  if (!state.token) return;

  group('Lessons - List', () => {
    const res = http.get(`${BASE_URL}/api/lessons?page=1&page_size=20`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    lessonsDuration.add(res.timings.duration);

    if (checkResponse('List Lessons', res)) {
      try {
        const body = JSON.parse(res.body);
        if (body.data && body.data.length > 0) {
          state.lessonIds = body.data.map(l => l.id);
        }
      } catch (e) {
        console.error('Parse error:', e);
      }
    }
  });

  if (state.lessonIds.length > 0) {
    group('Lessons - Get One', () => {
      const lId = randomItem(state.lessonIds);
      const res = http.get(`${BASE_URL}/api/lessons/${lId}`, {
        headers: { 'Authorization': `Bearer ${state.token}` },
      });
      lessonsDuration.add(res.timings.duration);
      checkResponse('Get Lesson', res);
    });
  }
}

function scenarioQuiz(state) {
  if (!state.token) return;

  group('Quiz - Get', () => {
    const res = http.get(`${BASE_URL}/api/quiz?count=10&license=PPL`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    quizDuration.add(res.timings.duration);
    checkResponse('Get Quiz', res);
  });
}

function scenarioHistory(state) {
  if (!state.token) return;

  group('History - Get', () => {
    const res = http.get(`${BASE_URL}/api/history?page=1&page_size=20`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    checkResponse('Get History', res);
  });

  group('Stats - Get', () => {
    const res = http.get(`${BASE_URL}/api/stats`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    checkResponse('Get Stats', res);
  });
}

function scenarioRecommendations(state) {
  if (!state.token) return;

  group('Recommendations - Get', () => {
    const res = http.get(`${BASE_URL}/api/recommendations`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    recommendationsDuration.add(res.timings.duration);
    checkResponse('Get Recommendations', res);
  });
}

function scenarioAdmin(state) {
  if (!state.token) return;

  group('Admin - Stats', () => {
    const res = http.get(`${BASE_URL}/api/admin/stats`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    adminDuration.add(res.timings.duration);
    checkResponse('Admin Stats', res);
  });

  group('Admin - Students', () => {
    const res = http.get(`${BASE_URL}/api/admin/students?page=1&page_size=10`, {
      headers: { 'Authorization': `Bearer ${state.token}` },
    });
    adminDuration.add(res.timings.duration);
    checkResponse('Admin Students', res);
  });
}

// ======================== SCÉNARIO PRINCIPAL ========================

export default function () {
  const state = getOrInitState(__VU);

  // Phase 1: Auth (toujours)
  scenarioAuth(state);

  // Phase 2: Parcours utilisateur aléatoire
  const scenario = Math.random();

  if (scenario < 0.25) {
    // 25% : Parcours questions
    scenarioQuestions(state);
    scenarioQuiz(state);
  } else if (scenario < 0.45) {
    // 20% : Parcours leçons
    scenarioLessons(state);
  } else if (scenario < 0.65) {
    // 20% : Parcours historique + stats
    scenarioHistory(state);
    scenarioRecommendations(state);
  } else if (scenario < 0.80) {
    // 15% : Parcours complet
    scenarioQuestions(state);
    scenarioLessons(state);
    scenarioQuiz(state);
    scenarioHistory(state);
    scenarioRecommendations(state);
  } else {
    // 20% : Parcours admin
    scenarioAdmin(state);
    scenarioQuestions(state);
    scenarioLessons(state);
  }

  // Pause entre les itérations
  sleep(Math.random() * 3 + 1);
}

// ======================== TEST DE SMOKE ========================

export function smokeTest() {
  // Test rapide pour vérifier que tout fonctionne
  const state = { token: null, studentId: null, questionIds: [], lessonIds: [] };

  scenarioAuth(state);
  if (state.token) {
    scenarioQuestions(state);
    scenarioLessons(state);
    scenarioQuiz(state);
    scenarioHistory(state);
    scenarioRecommendations(state);
    scenarioAdmin(state);
  }
}
