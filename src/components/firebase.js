// firebase.js
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyBno7yASW_08deDjyWj8AlZb5mVSmjpNyw",
  authDomain: "cards-app-8931e.firebaseapp.com",
  projectId: "cards-app-8931e",
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
