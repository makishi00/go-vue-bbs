<template>
    <form class="member-registration-form" action="">
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
                <button class="button is-primary" @click.prevent="memberRegistration">会員登録</button>
            </footer>
        </div>
    </form>
</template>

<script>
    import user from"../../service/user";

    export default {
        data() {
            return {
                email: "",
                password: ""
            }
        },
        methods:{
            async memberRegistration(){
                const response = await user.register(this.email, this.password);
                alert(await response.message);
                if (response.status === 200) {
                    location.href = '/login';
                }
            }
        }
    }
</script>