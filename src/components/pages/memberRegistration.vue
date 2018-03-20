<template>
    <div id="member-registration">
        <SideBar/>
        <MemberRegistration/>
    </div>
</template>

<script>
    import SideBar from '../modules/sidebar.vue';
    import MemberRegistration from '../modules/memberRegistrationForm.vue';

    export default {
        components: {
            SideBar,
            MemberRegistration
        }
    }
</script>

<style>
    div#member-registration {
        overflow: hidden;
    }
    div#member-registration form {
        width: 80%;
        margin-left: 17.5%;
        margin-top: 15%;
    }
</style>
