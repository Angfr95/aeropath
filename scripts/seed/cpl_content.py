"""Generate CPL chapters and their SQL import."""
import sys
sys.dont_write_bytecode = True
from cpl.mass_and_balance import build
from cpl import performance, principles_of_flight, flight_planning, communications, human_performance
from shared.content import generate
from shared.catalogue import write_catalogue
from cpl import instrumentation


def main():
    build()
    performance.build()
    principles_of_flight.build()
    flight_planning.build()
    communications.build()
    human_performance.build()
    instrumentation.build()
    generate('CPL')
    write_catalogue('CPL')


if __name__ == '__main__':
    main()
