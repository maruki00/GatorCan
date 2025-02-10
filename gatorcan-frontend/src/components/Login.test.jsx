import { render, screen, fireEvent } from "@testing-library/react";
import Login from "./Login"; // Adjust the import path as needed
import { MemoryRouter } from "react-router-dom";

describe("Login Component", () => {
  test("renders the login form with required fields and buttons", () => {
    render(
      <MemoryRouter>
        <Login />
      </MemoryRouter>
    );

    // Check if the login heading is displayed
    expect(screen.getByText(/login/i)).toBeInTheDocument();

    // Check if the username input field is rendered
    expect(screen.getByPlaceholderText(/username/i)).toBeInTheDocument();

    // Check if the password input field is rendered
    expect(screen.getByPlaceholderText(/password/i)).toBeInTheDocument();

    // Check if the login button is rendered
    expect(screen.getByRole("button", { name: /login/i })).toBeInTheDocument();

    // Check if the "Forgot Password?" link is rendered
    expect(screen.getByText(/forgot password\?/i)).toBeInTheDocument();
  });

  test("can type in the username and password fields", () => {
    render(
      <MemoryRouter>
        <Login />
      </MemoryRouter>
    );

    // Type in the username field
    fireEvent.change(screen.getByPlaceholderText(/username/i), {
      target: { value: "testuser" },
    });

    // Type in the password field
    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: "testpassword" },
    });

    // Check if the input fields have the correct values
    expect(screen.getByPlaceholderText(/username/i).value).toBe("testuser");
    expect(screen.getByPlaceholderText(/password/i).value).toBe("testpassword");
  });

  test("displays an error message when there is an error", () => {
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

    // Check if the error message appears
    expect(screen.getByRole("alert")).toHaveTextContent(/login failed/i);
  });

  test("focuses on username input on render", () => {
    render(
      <MemoryRouter>
        <Login />
      </MemoryRouter>
    );

    // Check if the username input is focused
    expect(screen.getByPlaceholderText(/username/i)).toHaveFocus();
  });
});
