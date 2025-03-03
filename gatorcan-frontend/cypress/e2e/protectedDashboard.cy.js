describe("ProtectedDashboard Tests", () => {
  beforeEach(() => {
    cy.clearLocalStorage();
  });

  it("Redirects to login if no user is logged in", () => {
    cy.visit("/dashboard");
    cy.url().should("include", "/login");
  });

  it("Redirects to dashboard if user lacks required roles", () => {
    cy.visit("/dashboard", {
      onBeforeLoad: (win) => {
        win.localStorage.setItem("username", "testuser");
        win.localStorage.setItem("roles", JSON.stringify(["guest"]));
      },
    });
    cy.url().should("include", "/dashboard");
  });

  it("Allows access if user has the correct role", () => {
    cy.visit("/dashboard", {
      onBeforeLoad: (win) => {
        win.localStorage.setItem("username", "testuser");
        win.localStorage.setItem("roles", JSON.stringify(["admin"]));
      },
    });

    cy.contains("Dashboard Content").should("be.visible");
  });
});
