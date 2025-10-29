import React from "react";
import './Login.css'
import handleLogin from "./Loginauth";
import handleSignup from "./SignupAuth";

export default function Login(props) {

    const handleLoginBtn = async () => {
    const username = document.querySelector('.username-input').value;
    const password = document.querySelector('.password-input').value;

    const [isLogged , user_id, token] = await handleLogin(username, password);

    if (isLogged) {
      props.setToken(token)  
      props.setUserId(user_id)
      props.changeFlag(true);
    } else {
      alert("Login failed");
    }
  };

  const handleSingUpBtn = async () => {
    const username = document.querySelector('.username-input').value;
    const password = document.querySelector('.password-input').value;

    const [isLogged , user_id, token] = await handleSignup(username, password);

    if (isLogged) {
      props.setToken(token)  
      props.setUserId(user_id)
      props.changeFlag(true);
    } else {
      alert("Login failed");
    }
  };


    return (
        <div className="main-login">
            <div className="user-details-div">
                <p className="tagline-text">One Card at a time</p>
                <p className="signin-text">Sign in with email</p>
                <div className="input-details-div">
                    <label className="title-label">
                        <input placeholder="username" type="text" className="username-input"></input>
                    </label>

                    <label className="description-label">
                        <input placeholder="password" type="password" className="password-input"></input>
                    </label>
                    <p className="forgot-password-text">Forgot Password?</p>
                    <button onClick={handleLoginBtn} className="login-btn">Login</button>
                    <button onClick={handleSingUpBtn} className="login-btn">Sign up</button>
                </div>


            </div>
        </div>
    )
}