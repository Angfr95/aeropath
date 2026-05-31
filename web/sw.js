const CACHE_NAME = "aeropath-v2";
const API_CACHE = "aeropath-api-v2";
const STATIC_CACHE = "aeropath-static-v2";

// Assets à pré-cacher au moment de l'installation
const PRECACHE_URLS = [
  "/",
  "/index.html",
  "/manifest.json",
  "/app.js",
  // "/icons/icon-192.png",  // Décommenter quand les PNG seront générés
  // "/icons/icon-512.png",  // Décommenter quand les PNG seront générés
];

// ======================== INSTALLATION ========================

self.addEventListener("install", (event) => {
  event.waitUntil(
    caches.open(STATIC_CACHE).then((cache) => {
      return cache.addAll(PRECACHE_URLS);
    })
  );
  self.skipWaiting();
});

// ======================== ACTIVATION ========================

self.addEventListener("activate", (event) => {
  event.waitUntil(
    caches.keys().then((cacheNames) => {
      return Promise.all(
        cacheNames
          .filter((name) => {
            return name !== CACHE_NAME && name !== API_CACHE && name !== STATIC_CACHE;
          })
          .map((name) => caches.delete(name))
      );
    })
  );
  self.clients.claim();
});

// ======================== STRATÉGIES DE CACHE ========================

// Stratégie: Network First avec fallback cache (pour l'API)
// ⚠️ Cache API ne supporte que les requêtes GET
async function networkFirst(request) {
  try {
    const response = await fetch(request);
    if (response.ok && request.method === "GET") {
      const cache = await caches.open(API_CACHE);
      cache.put(request, response.clone());
    }
    return response;
  } catch (error) {
    const cached = await caches.match(request);
    if (cached) {
      return cached;
    }
    // Fallback: retourner une réponse offline
    return new Response(
      JSON.stringify({ error: "offline", message: "Vous êtes hors-ligne" }),
      { status: 503, headers: { "Content-Type": "application/json" } }
    );
  }
}

// Stratégie: Cache First (pour les assets statiques)
async function cacheFirst(request) {
  const cached = await caches.match(request);
  if (cached) {
    return cached;
  }
  try {
    const response = await fetch(request);
    if (response.ok) {
      const cache = await caches.open(STATIC_CACHE);
      cache.put(request, response.clone());
    }
    return response;
  } catch (error) {
    return new Response("Offline", { status: 503 });
  }
}

// Stratégie: Stale While Revalidate (pour les pages)
async function staleWhileRevalidate(request) {
  const cache = await caches.open(CACHE_NAME);
  const cached = await cache.match(request);

  const fetchPromise = fetch(request)
    .then((response) => {
      if (response.ok) {
        cache.put(request, response.clone());
      }
      return response;
    })
    .catch(() => cached);

  return cached || fetchPromise;
}

// ======================== INTERCEPTION DES REQUÊTES ========================

self.addEventListener("fetch", (event) => {
  const { request } = event;

  // Ignorer les requêtes non-HTTP (chrome-extension, data, blob, etc.)
  if (!request.url.startsWith("http")) {
    return;
  }

  const url = new URL(request.url);

  // API requests → Network First
  if (url.pathname.startsWith("/api/") || url.pathname.startsWith("/auth/")) {
    event.respondWith(networkFirst(request));
    return;
  }

  // Swagger docs → Network First
  if (url.pathname.startsWith("/swagger/")) {
    event.respondWith(networkFirst(request));
    return;
  }

  // Assets statiques (CSS, JS, images) → Cache First
  if (
    request.destination === "style" ||
    request.destination === "script" ||
    request.destination === "image" ||
    request.destination === "font"
  ) {
    event.respondWith(cacheFirst(request));
    return;
  }

  // Pages HTML → Stale While Revalidate
  if (request.mode === "navigate") {
    event.respondWith(staleWhileRevalidate(request));
    return;
  }

  // Tout le reste → Network First
  event.respondWith(networkFirst(request));
});

// ======================== BACKGROUND SYNC ========================

self.addEventListener("sync", (event) => {
  if (event.tag === "sync-answers") {
    event.waitUntil(syncOfflineAnswers());
  }
});

async function syncOfflineAnswers() {
  try {
    const db = await openIndexedDB();
    const pendingAnswers = await db.getAll("pending_answers");

    for (const answer of pendingAnswers) {
      try {
        const response = await fetch("/api/questions/answer", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${answer.token}`,
          },
          body: JSON.stringify({
            question_id: answer.question_id,
            answer: answer.answer,
          }),
        });

        if (response.ok) {
          await db.delete("pending_answers", answer.id);
        }
      } catch (err) {
        console.error("[sync] erreur synchronisation:", err);
      }
    }
  } catch (err) {
    console.error("[sync] erreur base de données:", err);
  }
}

// ======================== PUSH NOTIFICATIONS ========================

self.addEventListener("push", (event) => {
  let data = { title: "AeroPath", body: "Nouvelle notification" };

  if (event.data) {
    try {
      data = event.data.json();
    } catch {
      data.body = event.data.text();
    }
  }

  const options = {
    body: data.body,
    // icon: "/icons/icon-192.png",  // Décommenter quand les PNG seront générés
    // badge: "/icons/icon-192.png", // Décommenter quand les PNG seront générés
    vibrate: [200, 100, 200],
    data: {
      url: data.url || "/",
    },
  };

  event.waitUntil(self.registration.showNotification(data.title, options));
});

self.addEventListener("notificationclick", (event) => {
  event.notification.close();
  const url = event.notification.data?.url || "/";
  event.waitUntil(clients.openWindow(url));
});

// ======================== INDEXED DB HELPER ========================

function openIndexedDB() {
  return new Promise((resolve, reject) => {
    const request = indexedDB.open("AeroPathOffline", 1);

    request.onupgradeneeded = (event) => {
      const db = event.target.result;
      if (!db.objectStoreNames.contains("pending_answers")) {
        db.createObjectStore("pending_answers", {
          keyPath: "id",
          autoIncrement: true,
        });
      }
      if (!db.objectStoreNames.contains("cached_questions")) {
        db.createObjectStore("cached_questions", {
          keyPath: "id",
        });
      }
    };

    request.onsuccess = (event) => {
      const db = event.target.result;

      resolve({
        getAll(storeName) {
          return new Promise((resolve, reject) => {
            const tx = db.transaction(storeName, "readonly");
            const store = tx.objectStore(storeName);
            const req = store.getAll();
            req.onsuccess = () => resolve(req.result);
            req.onerror = () => reject(req.error);
          });
        },
        put(storeName, value) {
          return new Promise((resolve, reject) => {
            const tx = db.transaction(storeName, "readwrite");
            const store = tx.objectStore(storeName);
            const req = store.put(value);
            req.onsuccess = () => resolve(req.result);
            req.onerror = () => reject(req.error);
          });
        },
        delete(storeName, key) {
          return new Promise((resolve, reject) => {
            const tx = db.transaction(storeName, "readwrite");
            const store = tx.objectStore(storeName);
            const req = store.delete(key);
            req.onsuccess = () => resolve();
            req.onerror = () => reject(req.error);
          });
        },
      });
    };

    request.onerror = () => reject(request.error);
  });
}
