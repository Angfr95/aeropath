"""Generate IR chapters and their SQL import."""
import sys
sys.dont_write_bytecode = True
from ir.navigation import build
from ir import approaches, instrumentation, meteorology, communications, flight_planning, air_law, human_performance
from shared.content import generate
from shared.catalogue import write_catalogue


def main():
    build()
    approaches.build()
    instrumentation.build()
    meteorology.build()
    communications.build()
    flight_planning.build()
    air_law.build()
    human_performance.build()
    generate('IR')
    write_catalogue('IR')


if __name__ == '__main__':
    main()
