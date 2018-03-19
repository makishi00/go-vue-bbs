<template>
    <login-form></login-form>
</template>

<script>
    import SideBar from '../modules/sidebar.vue';
    import LoginForm from '../modules/loginform.vue';

    export default {
        name: "login",
        components: {
            LoginForm,
            "side-bar": SideBar,
            "login-form": LoginForm
        }
    }
</script>