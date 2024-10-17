<script lang="ts">
	let count = $state(0);
	let countIncrement = $derived(count + 1);
</script>

<button onclick={() => count++}>
	Clicked {countIncrement} times
</button>
