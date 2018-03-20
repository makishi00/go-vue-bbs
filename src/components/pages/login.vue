<template>
    <div id="login">
        <SideBar/>
        <LoginForm/>
    </div>
</template>

<script>
    import SideBar from '../modules/sidebar.vue';
    import LoginForm from '../modules/loginForm.vue';

    export default {
        components: {
            SideBar,
            LoginForm
        }
    }
</script>

<style>
div#login {
    overflow: hidden;
}
div#login form {
    width: 80%;
    margin-left: 17.5%;
    margin-top: 15%;
}
</style>
