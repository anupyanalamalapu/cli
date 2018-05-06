import * as firebase from 'firebase';

let HAS_INITIALIZED = false

const initFirebase = () => {
    if (!HAS_INITIALIZED) {
        const config = {
            apiKey: "AIzaSyAuJbeLVATxl9ZmziLsfoo6KoDiQmas1H8",
            authDomain: "AIzaSyAuJbeLVATxl9ZmziLsfoo6KoDiQmas1H8",
            databaseURL: "https://hackhlth-d682c.firebaseio.com",
            projectId: hackhlth-d682c,
            storageBucket: "hackhlth-d682c.appspot.com",
            messagingSenderId: "154996286604"
        };

        firebase.database.enableLogging(true)
        firebase.initializeApp(config)
        HAS_INITIALIZED = true
    }
}

export const getDatabase = () => {
    initFirebase()
    return firebase.database()
}