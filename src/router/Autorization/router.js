import VueRouter from "vue-router";

// Autorization routings
import SignIn from "/src/pages/Autorization/components/SignIn.vue";
import SignUp from "/src/pages/Autorization/components/SignUp.vue";

export default new VueRouter({
    routes: [
        {
            path: "/signin",
            component: SignIn,
        },
        {
            path: "/signup",
            component: SignUp,
        },
    ],
    mode: "history",
});
