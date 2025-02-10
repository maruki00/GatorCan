import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import {
  Button,
  FormControl,
  FormGroup,
  FormControlLabel,
  Checkbox,
} from "@mui/material";

import { useRef, useState, useEffect } from "react";
import { Box } from "@mui/material";
import addUser from "../services/AdminService";

const UserRegistration = () => {
  const userRef = useRef();
  const errRef = useRef();

  const [user, setUser] = useState("");
  const [pwd, setPwd] = useState("");
  const [email, setEmail] = useState("");
  const [repwd, setRepwd] = useState("");
  const [errMsg, setErrMsg] = useState("");
  const [displaySuccess, setDisplaySuccess] = useState("");

  // State for roles
  const [selectedRoles, setSelectedRoles] = useState({
    student: false,
    admin: false,
    instructor: false,
  });

  useEffect(() => {
    userRef.current.focus();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [user, pwd, repwd, email, selectedRoles]);

  // Handle checkbox change for roles
  const handleRoleChange = (event) => {
    setSelectedRoles({
      ...selectedRoles,
      [event.target.name]: event.target.checked,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Validations
    if (pwd !== repwd) {
      setErrMsg("Passwords do not match");
      return;
    }

    // Get selected roles
    const roles = Object.keys(selectedRoles).filter(
      (role) => selectedRoles[role]
    );

    if (roles.length === 0) {
      setErrMsg("Please select at least one role");
      return;
    }

    try {
      // Simulate successful user creation
      const response = await addUser(user, pwd, email, roles);
      console.log("Add user API Successful:", response);
      let success = response["success"];
      if (!success) {
        console.log(response["message"]);
        setErrMsg(response["message"]);
      } else {
        setDisplaySuccess(
          "User with name: " + user + " created successfully!!"
        );
        setUser("");
        setEmail("");
        setPwd("");
        setRepwd("");
        setSelectedRoles({
          student: false,
          admin: false,
          instructor: false,
        });
      }
    } catch (error) {
      setErrMsg(error.response?.data?.message || "Unknown error");
    }
  };

  const paperStyle = {
    padding: 20,
    width: 300,
    margin: "19px auto",
  };
  const btnstyle = { backgroundColor: "#1B6DA1", margin: "20px 0" };
  const inputStyle = { margin: "12px auto" };
  const errorStyle = { color: "red" };

  return (
    <Box
      sx={{
        backgroundImage: `url('/Gator.png')`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Grid>
        {displaySuccess !== "" ? (
          <Paper elevation={12} style={paperStyle}>
            <Grid align="center">
              <p>{displaySuccess}</p>
            </Grid>
            <Button
              style={btnstyle}
              color="primary"
              variant="contained"
              fullWidth
              onClick={() => {
                setDisplaySuccess("");
              }}
            >
              Create new user
            </Button>
          </Paper>
        ) : (
          <div>
            <form onSubmit={handleSubmit}>
              <Paper elevation={12} style={paperStyle}>
                <Grid align="center">
                  <h2>Add user</h2>
                </Grid>
                <FormControl component="fieldset">
                  <FormGroup>
                    <FormControlLabel
                      control={
                        <Checkbox
                          checked={selectedRoles.student}
                          onChange={handleRoleChange}
                          name="student"
                        />
                      }
                      label="Student"
                    />
                    <FormControlLabel
                      control={
                        <Checkbox
                          checked={selectedRoles.admin}
                          onChange={handleRoleChange}
                          name="admin"
                        />
                      }
                      label="Admin"
                    />
                    <FormControlLabel
                      control={
                        <Checkbox
                          checked={selectedRoles.instructor}
                          onChange={handleRoleChange}
                          name="instructor"
                        />
                      }
                      label="Instructor"
                    />
                  </FormGroup>
                </FormControl>
                <Input
                  type="text"
                  id="username"
                  ref={userRef}
                  autoComplete="off"
                  onChange={(e) => setUser(e.target.value)}
                  value={user}
                  required
                  style={inputStyle}
                  placeholder="Username"
                  fullWidth
                />
                <Input
                  type="email"
                  id="email"
                  autoComplete="off"
                  onChange={(e) => setEmail(e.target.value)}
                  value={email}
                  required
                  style={inputStyle}
                  placeholder="Email"
                  fullWidth
                />
                <Input
                  type="password"
                  id="password"
                  onChange={(e) => setPwd(e.target.value)}
                  value={pwd}
                  required
                  placeholder="Password"
                  style={inputStyle}
                  fullWidth
                />
                <Input
                  type="password"
                  id="repassword"
                  onChange={(e) => setRepwd(e.target.value)}
                  value={repwd}
                  required
                  placeholder="Re-enter Password"
                  style={inputStyle}
                  fullWidth
                />
                <Button
                  style={btnstyle}
                  type="submit"
                  color="primary"
                  variant="contained"
                  fullWidth
                >
                  Submit
                </Button>
              </Paper>
            </form>
            <p
              ref={errRef}
              className={errMsg ? "errmsg" : "offscreen"}
              aria-live="assertive"
              style={errorStyle}
            >
              {errMsg}
            </p>
          </div>
        )}
      </Grid>
    </Box>
  );
};

export default UserRegistration;
