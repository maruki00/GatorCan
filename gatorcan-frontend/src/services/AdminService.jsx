import axios from "axios";

const base_url = "/";

export const addUser = async (username, password, email, roles) => {
  const add_user_url = base_url + "admin/add_user";
  console.log(roles);
  var roles_string = "";
  roles.forEach((role) => {
    roles_string = roles_string + '"' + role + '",';
  });
  roles_string = roles_string.slice(0, -1);
  console.log(roles_string);

  try {
    const refreshToken = localStorage.getItem("refreshToken");
    const response = await axios.post(
      add_user_url,
      { username, password, email, roles },
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + refreshToken,
        },
      }
    );
    console.log("RESPONSE RECEIVED: ", response.data);
    return { success: true };
  } catch (err) {
    if (err.response) {
      console.error("Add user failed:", err.response.data);
      return {
        success: false,
        message: "Add user failed: " + (err.response.data?.error || ""),
      };
    } else if (err.request) {
      console.error("No response received from server: ", err.request);
      return { success: false, message: "No response received from server" };
    } else {
      console.error("AXIOS ERROR:", err.message);
      return { success: false, message: "Unknown error" };
    }
  }
};

export default addUser;
