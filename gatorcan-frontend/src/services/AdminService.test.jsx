import axios from "axios";
import { addUser } from "./AdminService";

jest.mock("axios");

describe("AdminService - addUser", () => {
  test("successfully adds a user", async () => {
    axios.post.mockResolvedValue({
      data: { message: "User added successfully" },
    });

    const result = await addUser(
      "testuser",
      "password123",
      "test@example.com",
      ["admin"]
    );

    expect(result).toEqual({ success: true });
    expect(axios.post).toHaveBeenCalledWith(
      "/admin/add_user",
      {
        username: "testuser",
        password: "password123",
        email: "test@example.com",
        roles: ["admin"],
      },
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
      })
    );
  });
});
