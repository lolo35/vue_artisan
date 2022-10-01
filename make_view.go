package main

import (
	"fmt"
)

func (a Artisan) MakeView(arguments []string) {
	if len(arguments) == 1 {
		fmt.Println(red + "make:view expects view name to be provided")
		return
	}

	fmt.Println("Creating view...")

	fileContents := `<script setup lang="ts">
	
</script>
<template>
	<div>
	
	</div>
</template>
<style scoped>

</style>`

	a.CreateFile("views/", arguments[1]+".vue", fileContents)

}
