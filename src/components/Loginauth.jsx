import { signInWithEmailAndPassword } from "firebase/auth";
import { auth } from "./firebase";


export default async function handleLogin(email, password) {
    try {
        const userCredential = await signInWithEmailAndPassword(auth, email, password);
        const user = userCredential.user;
        const token = await user.getIdToken();
        const user_id = await user.uid;
        return [true, user_id, token]


    } catch (error) {
        console.error("Login failed:", error.code, error.message);

        if (error.code === "auth/wrong-password") {
            alert("Incorrect password");
        } else if (error.code === "auth/user-not-found") {
            alert("No user with this email");
        } else if (error.code === "auth/invalid-email") {
            alert("Invalid email format");
        } else {
            alert("Login failed");
        }

         return [false, "Failed"];
    }

}
