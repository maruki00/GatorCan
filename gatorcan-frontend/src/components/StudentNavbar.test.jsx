import React from "react";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import StudentNavbar from "./StudentNavbar";
import "@testing-library/jest-dom";

describe("StudentNavbar Component", () => {
  test("renders all navbar items", () => {
    render(
      <MemoryRouter>
        <StudentNavbar />
      </MemoryRouter>
    );

    // Check for all navbar text labels
    expect(screen.getByText("GatorCan")).toBeInTheDocument();
    expect(screen.getByText("Profile")).toBeInTheDocument();
    expect(screen.getByText("Dashboard")).toBeInTheDocument();
    expect(screen.getByText("Courses")).toBeInTheDocument();
    expect(screen.getByText("Calendar")).toBeInTheDocument();
    expect(screen.getByText("Inbox")).toBeInTheDocument();
    expect(screen.getByText("Logout")).toBeInTheDocument();
  });
});
