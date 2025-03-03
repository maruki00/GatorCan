describe('Login Page Tests', () => {
  beforeEach(() => {
    cy.visit('http://localhost:5173'); // Update with actual frontend URL if needed
  });

  it('should display the login page with all required components', () => {
    cy.get('[data-testid="cypress-title"]').should('have.text', 'Login');
    cy.get('#username').should('exist');
    cy.get('#password').should('exist');
    cy.get('button[type="submit"]').should('exist').and('have.text', 'Login');
    cy.get('a').contains('Forgot Password?').should('exist');
  });
  it('should allow a user to enter credentials and attempt login', () => {
    cy.get('#username').type('admin');
    cy.get('#password').type('admin');
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/admin-dashboard'); // Update this based on actual redirection behavior

    // Optionally, check localStorage to see if login was successful
    cy.window().its('localStorage.username').should('exist');
  });
});