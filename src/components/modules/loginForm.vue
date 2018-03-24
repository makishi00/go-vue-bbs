<template>
    <form class="login-form" action="">
        <div class="modal-card" style="width: auto">
            <header class="modal-card-head">
                <p class="modal-card-title">ログイン</p>
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
                <button class="button is-primary" @click.prevent="login">ログイン</button>
            </footer>
        </div>
    </form>
</template>

<script>
    import auth from"../../service/auth";

    export default {
        data() {
            return {
                email: "",
                password: ""
            }
        },
        methods:{
            async login(){
                const response = await auth.login(this.email, this.password);
                alert(await response.message);
                if (response.status === 200) {
                    auth.setLoginToken(await response.data.token);
                    location.href = '/';
                }
            }
        }
    }
</script>