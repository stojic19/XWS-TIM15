import firebase from 'firebase/compat/app';
import 'firebase/compat/storage';
import 'firebase/compat/firestore';

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyAeIv5Mgim1eAedFr6spAZxncsqZcDW8Oo",
    authDomain: "dislinkt-1601a.firebaseapp.com",
    projectId: "dislinkt-1601a",
    storageBucket: "dislinkt-1601a.appspot.com",
    messagingSenderId: "921298308920",
    appId: "1:921298308920:web:07090474ca84dc9477fbcd",
    measurementId: "G-PTLT6XJXME"
  };
  
// Initialize Firebase
firebase.initializeApp(firebaseConfig);

const projectStorage = firebase.storage();
const projectFirestore = firebase.firestore();
const timestamp = firebase.firestore.FieldValue.serverTimestamp;

export { projectStorage, projectFirestore, timestamp };