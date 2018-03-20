<template>
    <div id="login">
        <SideBar/>
        <LoginForm/>
    </div>
</template>

<script>
    import SideBar from '../modules/sidebar.vue';
    import LoginForm from '../modules/loginform.vue';

    export default {
        name: "login",
        components: {
            SideBar,
            LoginForm
        }
    }
</script>

