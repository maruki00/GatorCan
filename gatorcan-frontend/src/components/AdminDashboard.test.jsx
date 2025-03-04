import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { MemoryRouter, useNavigate } from "react-router-dom";
import { describe, test, expect, jest } from "@jest/globals";
import AdminDashboard from "./AdminDashboard";
import "@testing-library/jest-dom";

describe("Admin Dashboard Component", () => {
    test("add user renders correctly", () => {
        render(
            <MemoryRouter>
                <AdminDashboard></AdminDashboard>
            </MemoryRouter>
        )
        expect(screen.getByText(/add user/i)).toBeInTheDocument();

    });
});