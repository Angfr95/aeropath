"""Generate the PPL SQL import; retain the existing helper import interface."""
from shared.content import CATEGORIES, ROOT, generate, sql, stable_id, upsert

if __name__ == '__main__':
    generate('PPL')
