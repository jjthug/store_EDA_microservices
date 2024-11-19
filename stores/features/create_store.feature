Feature: Create Store

    As a store owner i should be able to add new stores

    Scenario: Creating a store called "XYZ"
        Given a valid store owner
        And no store called "XYZ" exists
        When i create the store called "XYZ"
        Then a store called "XYZ" exists