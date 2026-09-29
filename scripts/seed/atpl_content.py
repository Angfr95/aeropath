"""Generate ATPL chapters and their SQL import."""
import sys
sys.dont_write_bytecode = True
from atpl.flight_planning import build
from atpl import aircraft_general, instrumentation, navigation, communications, pressurisation
from shared.content import generate
from shared.catalogue import write_catalogue
from atpl import mass_and_balance


def main():
    build()
    aircraft_general.build()
    instrumentation.build()
    navigation.build()
    communications.build()
    pressurisation.build()
    mass_and_balance.build()
    generate('ATPL')
    write_catalogue('ATPL')


if __name__ == '__main__':
    main()
