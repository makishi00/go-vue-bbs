<template>
    <side-bar></side-bar>
</template>

<form action="">
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <p class="modal-card-title">会員登録</p>
        </header>
        <section class="modal-card-body">
            <b-field label="Email">
                <b-input
                        type="email"
                        v-model="email"
                        placeholder="Your email"
                        required>
                </b-input>
            </b-field>

            <b-field label="Password">
                <b-input
                        type="password"
                        v-model="password"
                        password-reveal
                        placeholder="Your password"
                        required>
                </b-input>
            </b-field>

            <b-checkbox>Remember me</b-checkbox>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="registration">登録</button>
        </footer>
    </div>
</form>

<script>
    import SideBar from '../modules/sidebar.vue';
    export default {
        name: "memberRegistration",
        components: {
            SideBar
        },
        data() {
            return {
                email: "",
                password: ""
            }
        },
        methods:{
            registration(){
                console.log("registrationBtn click")
            }
        }
    }
</script>