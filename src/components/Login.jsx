import React from "react";
import './Login.css'
import handleLogin from "./auth";

export default function Login(props) {

    const handleClick = async () => {
    const username = document.querySelector('.username-input').value;
    const password = document.querySelector('.password-input').value;

    const [isLogged , token] = await handleLogin(username, password); // 👈 wait for it to resolve

    if (isLogged) {
      props.setToken(token)  
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
                    <button onClick={handleClick} className="login-btn">Get Started</button>
                </div>


            </div>
        </div>
    )
}