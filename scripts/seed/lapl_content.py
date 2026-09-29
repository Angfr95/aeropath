"""Generate LAPL chapters and their SQL import."""
import sys
sys.dont_write_bytecode = True
from lapl.catalogue import build
from shared.content import generate
from shared.catalogue import write_catalogue


def main():
    build()
    generate('LAPL')
    write_catalogue('LAPL')


if __name__ == '__main__':
    main()
