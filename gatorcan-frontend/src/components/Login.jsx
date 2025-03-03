import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import { Button, Typography } from "@mui/material";

import { Link, replace, useNavigate, Navigate } from "react-router-dom";
import { useRef, useState, useEffect } from "react";
import { Box } from "@mui/material";
import login from "../services/AuthService";


const Login = () => {

  const username = localStorage.getItem("username");
  if (username) {
    return <Navigate to="/dashboard" replace />;
  }

  const userRef = useRef();
  const errRef = useRef();

  const navigate = useNavigate();

  const [user, setUser] = useState("");
  const [pwd, setPwd] = useState("");
  const [errMsg, setErrMsg] = useState("");

  useEffect(() => {
    userRef.current.focus();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [user, pwd]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await login(user, pwd);
      console.log("Login API Successful:", response);
      let success = response["success"];
      if (!success) {
        console.log(response["message"]);
        setErrMsg(response["message"])
      } else {
        navigate("/dashboard", replace);
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
  const inputStyle = { margin: "20px auto" };
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
        <form onSubmit={handleSubmit}>
          <Paper elavation={12} style={paperStyle}>
            <Grid align="center">
              <h2 data-testid="cypress-title">Login</h2>
            </Grid>
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
              type="password"
              id="password"
              onChange={(e) => setPwd(e.target.value)}
              value={pwd}
              required
              placeholder="Password"
              fullWidth
            />

            <Button
              style={btnstyle}
              type="submit"
              color="primary"
              variant="contained"
              fullWidth
            >
              Login
            </Button>
            <Typography>
              <Link to="/resetPassword">Forgot Password?</Link>
            </Typography>
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
      </Grid>
    </Box>
  );
};

export default Login;
