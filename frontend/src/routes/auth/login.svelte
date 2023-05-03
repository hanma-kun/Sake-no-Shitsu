<script lang="ts">
    import { LoginUser,SigninUser } from "../../utils/requests/post";
    import { useNavigate, useLocation } from "svelte-navigator";
    import { token,Alerts } from "../../store";
    import Alert from "../../utils/Alert.svelte";

    let lft = false;
    let navigate = useNavigate();
    let location = useLocation();
    $: if ($token) {
        navigate(`/`, { replace: true });
    }
    let loginusername: string;
    let loginpassword: string;

    let signemail: string;
    let signinpassword: string;
    let signinfirstname: string;
    let sigininlastname: string;
    let signinphoneno: string;

    Alerts.subscribe((value)=>{
        console.log(value)
    })

    function LoginFunc() {
        if (loginpassword != "" && loginusername != "") {
            LoginUser({
                email: loginusername,
                password: loginpassword,
            });
        }
    }
    

    function SignUpFunc(){
        if (sigininlastname!="" && signinfirstname !="" && signemail != "" && signinpassword!="" && signinphoneno !="" ){
            SigninUser({email:signemail,firstname:signinfirstname,lastname:sigininlastname,password:signinpassword,phone:signinpassword})
        }
    }

</script>

<div class="maincontainer">
    <div class="container" id="container" class:right-panel-active={lft}>
        <div class="form-container sign-up-container">
            <form action="#">
                <h1>Sign Up</h1>
                <input
                    type="text"
                    bind:value={signinfirstname}
                    placeholder="First Name"
                    required
                    />
                <input
                    type="text"
                    bind:value={sigininlastname}
                    placeholder="Last Name"
                    required
                    />
                <input
                    type="email"
                    bind:value={signemail}
                    placeholder="Email"
                    required
                    />
                <input
                    type="number"
                    bind:value={signinphoneno}
                    placeholder="Phone no"
                    required
                    />
                <input
                    type="password"
                    bind:value={signinpassword}
                    placeholder="Password"
                    required
                />
                <button type="submit" on:click|preventDefault={SignUpFunc}>Sign Up</button>
            </form>
        </div>
        <div class="form-container sign-in-container">
            <form action="#">
                <h1>Sign In</h1>
                <input
                    type="email"
                    bind:value={loginusername}
                    placeholder="Email"
                    required
                />
                <input
                    type="password"
                    bind:value={loginpassword}
                    placeholder="Password"
                    required
                />
                <button type="submit" on:click|preventDefault={LoginFunc}
                    >Sign In</button
                >
            </form>
        </div>

        <div class="overlay-container">
            <div class="overlay">
                <div class="overlay-panel overlay-right">
                    <h1>Welcome Back!</h1>
                    <p>
                        To stay connected with us please login with your
                        personal info
                    </p>
                    <button
                        class="ghost"
                        id="SignUp"
                        on:click={() => {
                            lft = true;
                        }}>New User ? Sign up</button
                    >
                </div>
                <div class="overlay-panel overlay-left">
                    <h1>Hello, Friend!</h1>

                    <p>
                        Enter your personal details and start your journey with
                        us
                    </p>
                    <button
                        class="ghost"
                        id="SignIn"
                        on:click={() => {
                            lft = false;
                        }}>Already there ? Sign In</button
                    >
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .maincontainer {
        width: 100%;
        height: 100%;
        overflow: hidden;
        display: grid;
        place-items: center;
    }

    p {
        font-size: 14px;
        font-weight: 100;
        line-height: 20px;
        letter-spacing: 0.5px;
        margin: 20px 0 30px;
    }

    button {
        border-radius: 10px;
        background-color: transparent;
        border: 1px solid rgba(0, 0, 0, 0.5);
        font-size: 12px;
        font-weight: bold;
        padding: 12px 45px;
        letter-spacing: 1px;
        text-transform: uppercase;
        transition: transform 80ms ease-in;
    }

    form {
        background-color: #ffffff;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        padding: 0 50px;
        height: 100%;
        text-align: center;
    }

    input {
        background-color: #fff;
        border: 1px solid rgba(0, 0, 0, 0.39);
        border-radius: 5px;
        padding: 12px 15px;
        margin: 8px 0;
        width: 100%;
    }

    .container {
        background-color: #f0f0f0;
        border-radius: 10px;
        box-shadow: 0 14px 28px rgba(0, 0, 0, 0.05),
            0 10px 10px rgba(0, 0, 0, 0.09);
        position: relative;
        overflow: hidden;
        width: 768px;
        min-height: 480px;
        margin-top: 30px;
    }
    .form-container {
        position: absolute;
        top: 0;
        height: 100%;
        transition: all 0.6s ease-in-out;
    }

    .sign-in-container {
        left: 0;
        width: 50%;
        z-index: 2;
    }

    .sign-up-container {
        left: 0;
        width: 50%;
        opacity: 0;
        z-index: 1;
    }

    .container.right-panel-active .sign-in-container {
        transform: translateX(100%);
    }

    .container.right-panel-active .sign-up-container {
        transform: translateX(100%);
        opacity: 1;
        z-index: 5;
        animation: show 0.6s;
    }

    @keyframes show {
        0%,
        49.99% {
            opacity: 0;
            z-index: 1;
        }

        50%,
        100% {
            opacity: 1;
            z-index: 5;
        }
    }
    .overlay-container {
        position: absolute;
        top: 0;
        left: 50%;
        width: 50%;
        height: 100%;
        overflow: hidden;
        transition: transform 0.6s ease-in-out;
        z-index: 100;
    }

    .container.right-panel-active .overlay-container {
        transform: translateX(-100%);
    }

    .overlay {
        background-repeat: no-repeat;
        background-size: cover;
        background-position: 0 0;
        color: black;
        position: relative;
        left: -100%;
        height: 100%;
        width: 200%;
        transform: translateX(0);
        transition: transform 0.6s ease-in-out;
    }

    .container.right-panel-active .overlay {
        transform: translateX(50%);
    }

    .overlay-panel {
        position: absolute;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        padding: 0 40px;
        text-align: center;
        top: 0;
        height: 100%;
        width: 50%;
        transform: translateX(0);
        transition: transform 0.6s ease-in-out;
    }

    .overlay-left {
        transform: translateX(-20%);
    }

    .container.right-panel-active .overlay-left {
        transform: translateX(0);
    }

    .overlay-right {
        right: 0;
        transform: translateX(0);
    }

    .container.right-panel-active .overlay-right {
        transform: translateX(20%);
    }
</style>
