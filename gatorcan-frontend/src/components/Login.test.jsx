import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import { describe, test, expect } from "@jest/globals";
import Login from "./Login";
import "@testing-library/jest-dom";

describe("Login Component", () => {
  test("renders the login form with required fields and buttons", () => {
    render(
      <MemoryRouter>
        <Login />
      </MemoryRouter>
    );

    expect(screen.getByRole("heading", { name: /login/i })).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /login/i })).toBeInTheDocument();
    expect(screen.getByText(/forgot password\?/i)).toBeInTheDocument();
  });

  test("can type in the username and password fields", () => {
    render(
      <MemoryRouter>
        <Login />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText(/username/i), {
      target: { value: "testuser" },
    });

    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: "testpassword" },
    });

    expect(screen.getByPlaceholderText(/username/i).value).toBe("testuser");
    expect(screen.getByPlaceholderText(/password/i).value).toBe("testpassword");
  });

  test("displays an error message when there is an error", async () => {
    render(
      <MemoryRouter>
        <Login />
      </MemoryRouter>
    );

    // Simulate setting an error message
    fireEvent.change(screen.getByPlaceholderText(/username/i), {
      target: { value: "invaliduser" },
    });
    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: "wrongpassword" },
    });

    // Trigger form submission to simulate an error
    fireEvent.click(screen.getByRole("button", { name: /login/i }));

    await waitFor(() => {
      const errorElement = screen.getByRole("alert"); // Selecting by role (more accessible)
      expect(errorElement).toBeInTheDocument();
      expect(errorElement).toHaveTextContent(/.+/); // Ensure it's not empty
      expect(errorElement).toHaveClass("errmsg"); // Ensure the correct class is applied     
    });
  });
});
