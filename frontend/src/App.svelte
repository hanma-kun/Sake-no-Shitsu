<script>
  import { Router, Route, Link } from "svelte-navigator";
  import Login from "./routes/auth/login.svelte";
  import PrivateRoute from "./utils/PrivateRoute.svelte";
  import { fade } from "svelte/transition";
  import Cart from "./routes/cart/cart.svelte";
  import Navigator from "./utils/navigator.svelte";
  import Home from "./routes/home/home.svelte";
  import Loader from "./utils/loader.svelte";
  import { onMount } from "svelte";
  import Alert from "./utils/Alert.svelte";
    import { VerifyToken } from "./utils/requests/get";
    import Product from "./routes/cart/product.svelte";
    import Preview from "./routes/cart/preview.svelte";
  let pageloaded = false;
  
  onMount(() =>  {
    setTimeout(async () => {
      pageloaded = true;
      await VerifyToken();
      window.scroll(0, 0);
    }, 2000);
  });
</script>

<Alert />
{#if !pageloaded}
  <div out:fade>
    <Loader />
  </div>
{/if}

<Router primary={false}>
  <Navigator />
  <main>
    <Route path="signin" component={Login} />
    <Route path="/" component={Home} />
    <Route path="/preview" component={Preview}/>
    <Route path="product" component={Product} />
    <PrivateRoute path="cart" let:location><Cart /></PrivateRoute>
  </main>
</Router>

<style>
  main {
    width: 100%;
    height: calc(100vh - 60px);
  }
</style>
