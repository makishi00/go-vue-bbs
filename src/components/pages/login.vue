<template>
<<<<<<< Updated upstream
    <div id="login">
        <SideBar/>
        <LoginForm/>
=======
    <div class="logForm">
        <side-bar style="float: left;"></side-bar>
        <div class="log">
            <login-form></login-form>
        </div>
>>>>>>> Stashed changes
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

<<<<<<< Updated upstream
=======
<style>
div.logForm {
    overflow: hidden;
}
div.log {
    width: 80%;
    margin-left: 2.5%;
    margin-top: 15%;
    float: left;
}
</style>
>>>>>>> Stashed changes
