import { signInWithEmailAndPassword } from "firebase/auth";
import { auth } from "./firebase";


export default async function handleLogin(email, password) {
    //   try {
    // const userCredential = await signInWithEmailAndPassword(auth, email, password);
    // console.log(userCredential)
    // const user = userCredential.user;
    // const token = await user.getIdToken();

    try {
        const userCredential = await signInWithEmailAndPassword(auth, email, password);
        const user = userCredential.user;
        const token = await user.getIdToken();

        return [true, token]


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



    //     const response = await fetch("http://localhost:8080/add/card", {
    //       method: "POST",
    //       headers: {
    //         "Content-Type": "application/json",
    //         "Authorization": `Bearer ${token}`,
    //       },
    //       body: JSON.stringify({ title: "test", description: "card" }),
    //     });

    //     if (!response.ok) {
    //       console.error("Backend responded with error:", response.status);
    //       return false; // ✅ explicitly return
    //     }

    //     return true; // ✅ explicitly return
    //   } catch (error) {
    //     console.error("Login failed:", error.message);
    //     return false; // ✅ explicitly return
    //   }
}
