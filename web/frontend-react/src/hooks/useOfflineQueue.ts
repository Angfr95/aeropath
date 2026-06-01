import { useState, useCallback } from "react";
import type { PendingAnswer } from "@/types";

// ---------------------------------------------------------------------------
// Petit wrapper IndexedDB offline en attendant la roadmap "sync"
// ---------------------------------------------------------------------------

const DB_NAME = "AeroPathOffline";
const DB_VERSION = 1;

function openDB(): Promise<IDBDatabase> {
  return new Promise((resolve, reject) => {
    const req = indexedDB.open(DB_NAME, DB_VERSION);
    req.onupgradeneeded = () => {
      const db = req.result;
      if (!db.objectStoreNames.contains("pending_answers")) {
        db.createObjectStore("pending_answers", { keyPath: "id", autoIncrement: true });
      }
    };
    req.onsuccess = () => resolve(req.result);
    req.onerror = () => reject(req.error);
  });
}

// ---------------------------------------------------------------------------

export function useOfflineQueue() {
  const [queueLength, setQueueLength] = useState(0);

  const enqueue = useCallback(async (answer: Omit<PendingAnswer, "id" | "timestamp">) => {
    const db = await openDB();
    const tx = db.transaction("pending_answers", "readwrite");
    tx.objectStore("pending_answers").add({ ...answer, timestamp: Date.now() });
    tx.oncomplete = () => setQueueLength((n) => n + 1);
  }, []);

  const flush = useCallback(async () => {
    const db = await openDB();
    const tx = db.transaction("pending_answers", "readonly");
    const all: PendingAnswer[] = await new Promise((res) => {
      const req = tx.objectStore("pending_answers").getAll();
      req.onsuccess = () => res(req.result);
    });
    // La sync effective sera faite par le SW (background sync)
    setQueueLength(all.length);
    return all;
  }, []);

  return { enqueue, flush, queueLength };
}