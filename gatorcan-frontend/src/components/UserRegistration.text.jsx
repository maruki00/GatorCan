import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import UserRegistration from "./UserRegistration";
import addUser from "../services/AdminService";
import "@testing-library/jest-dom";

jest.mock("../services/AdminService");

describe("UserRegistration Component", () => {
  test("renders form with all elements", () => {
    render(<UserRegistration />);

    expect(
      screen.getByRole("heading", { name: /add user/i })
    ).toBeInTheDocument();

    expect(screen.getByPlaceholderText(/username/i)).toBeInTheDocument();
    expect(screen.getByPlaceholderText(/email/i)).toBeInTheDocument();
    expect(screen.getByPlaceholderText(/password/i)).toBeInTheDocument();
    expect(
      screen.getByPlaceholderText(/re-enter password/i)
    ).toBeInTheDocument();

    expect(screen.getByLabelText(/student/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/admin/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/instructor/i)).toBeInTheDocument();

    expect(screen.getByRole("button", { name: /submit/i })).toBeInTheDocument();
  });

  test("displays error when passwords do not match", async () => {
    render(<UserRegistration />);

    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: "password123" },
    });
    fireEvent.change(screen.getByPlaceholderText(/re-enter password/i), {
      target: { value: "password456" },
    });

    fireEvent.click(screen.getByRole("button", { name: /submit/i }));

    expect(screen.getByText(/passwords do not match/i)).toBeInTheDocument();
  });

  test("displays error when no role is selected", async () => {
    render(<UserRegistration />);

    fireEvent.change(screen.getByPlaceholderText(/username/i), {
      target: { value: "testuser" },
    });
    fireEvent.change(screen.getByPlaceholderText(/email/i), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: "password123" },
    });
    fireEvent.change(screen.getByPlaceholderText(/re-enter password/i), {
      target: { value: "password123" },
    });

    fireEvent.click(screen.getByRole("button", { name: /submit/i }));

    expect(
      screen.getByText(/please select at least one role/i)
    ).toBeInTheDocument();
  });

  test("submits form successfully and displays success message", async () => {
    addUser.mockResolvedValue({ success: true }); // Mock successful API call

    render(<UserRegistration />);

    fireEvent.change(screen.getByPlaceholderText(/username/i), {
      target: { value: "testuser" },
    });
    fireEvent.change(screen.getByPlaceholderText(/email/i), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: "password123" },
    });
    fireEvent.change(screen.getByPlaceholderText(/re-enter password/i), {
      target: { value: "password123" },
    });

    fireEvent.click(screen.getByLabelText(/student/i));

    fireEvent.click(screen.getByRole("button", { name: /submit/i }));

    await waitFor(() =>
      expect(
        screen.getByText(/user with name: testuser created successfully/i)
      ).toBeInTheDocument()
    );

    expect(screen.getByPlaceholderText(/username/i)).toHaveValue("");
    expect(screen.getByPlaceholderText(/email/i)).toHaveValue("");
    expect(screen.getByPlaceholderText(/password/i)).toHaveValue("");
    expect(screen.getByPlaceholderText(/re-enter password/i)).toHaveValue("");
    expect(screen.getByLabelText(/student/i)).not.toBeChecked();
  });
});
