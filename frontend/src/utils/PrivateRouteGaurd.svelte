<script>
    import { useNavigate, useLocation } from "svelte-navigator";
    import { token } from "../store"
  
    const navigate = useNavigate();
    const location = useLocation();
    
    $: if (!$token) {
      navigate("/signin", {
        state: { from: $location.pathname },
        replace: true,
      });
    }
  </script>
  
  {#if $token}
    <slot />
  {/if}
  