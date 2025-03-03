import axios from "axios";
import { login } from "./AuthService";
import { jwtDecode } from "jwt-decode";

jest.mock("axios");
jest.mock("jwt-decode", () => ({
  jwtDecode: jest.fn(),
}));

describe("AuthService - login", () => {
  test("successfully logs in a user and stores token data", async () => {
    const fakeToken = "fake.jwt.token";
    axios.post.mockResolvedValue({ data: { token: fakeToken } });

    jwtDecode.mockReturnValue({
      username: "testuser",
      roles: ["admin"],
      iss: "auth-service",
      exp: 1712345678,
      iat: 1712341234,
    });

    const result = await login("testuser", "password123");

    expect(result).toEqual({ success: true });
    expect(axios.post).toHaveBeenCalledWith(
      "/login",
      { username: "testuser", password: "password123" },
      { headers: { "Content-Type": "application/json" } }
    );

    expect(localStorage.getItem("refreshToken")).toBe(fakeToken);
    expect(localStorage.getItem("username")).toBe("testuser");
    expect(localStorage.getItem("roles")).toBe(JSON.stringify(["admin"]));
    expect(localStorage.getItem("iss")).toBe("auth-service");
    expect(localStorage.getItem("exp")).toBe("1712345678");
    expect(localStorage.getItem("iat")).toBe("1712341234");
  });
});
