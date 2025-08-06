import React from "react";
import './Login.css'

export default function Login(props){

    const handleClick = () =>{
        props.changeFlag(true)
    }

    return(
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