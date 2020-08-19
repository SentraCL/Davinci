<template>
	<span class="input_check">
		<span class="label">{{label}}</span>
		<input :disabled="isInactive" type="radio" @click="setValue(1)" :id="idY" :name="label" value="1" :checked="value==1">
		<label :for="idY" @click="setValue(1)">Si</label>
		<input :disabled="isInactive" type="radio" @click="setValue(0)" :id="idN" :name="label" value="0" :checked="value==0">
		<label :for="idN" @click="setValue(0)">No</label>
	</span>
</template>
<script>
	export default {
		name: "switch-box",
		props: {
			label: String,
			value: {},
			isInactive:false
		},
		data: function () {
			var _id = this.makeId(8);
			this.$emit("update:value", !this.value?0:this.value);
			this.$emit("update");			
			return {
				idY: _id + "Y",
				idN: _id + "N",
			};
		},
		methods: {


			setValue(value) {
				this.$emit("update:value", value);
				this.$emit("update");
			}
		}
	};
</script>
<style scoped>
	.label {
		min-width: 120px;
		padding-right: 120px;

		position: relative;
		left: 10px
	}

	.input_check {
		top: 10px;
		left: 6px;
		bottom: 2px;
		background-color: #EADFCE;
		position: relative;
		z-index: 1;
		display: inline-block;
		border: 1px solid #4E1D00;
		border-radius: 4px;
		-webkit-transition: background-color 0.2s, border-color 0.2s;
		transition: background-color 0.2s, border-color 0.2s;
	}

	input[type=radio],
	input[type=checkbox] {
		display: none;
	}

	.input_field {
		position: relative;
		display: block;
		background: #EADFCE;
		color: #000;
		-webkit-appearance: none;
	}


	input[type=radio]+label,
	input[type=checkbox]+label {
		display: inline-block;
		margin: -2px;
		padding: 4px 12px;
		margin-bottom: 0;
		font-size: 14px;
		line-height: 20px;
		text-align: center;
		vertical-align: middle;
		cursor: pointer;
		filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#EADFCE', endColorstr='#ffe6e6e6', GradientType=0);
		filter: progid:DXImageTransform.Microsoft.gradient(enabled=false);
	}

	input[type=radio]:checked+label,
	input[type=checkbox]:checked+label {
		background-image: none;
		outline: 0;
		-webkit-box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.15), 0 1px 2px rgba(0, 0, 0, 0.05);
		-moz-box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.15), 0 1px 2px rgba(0, 0, 0, 0.05);
		box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.15), 0 1px 2px rgba(0, 0, 0, 0.05);
		background-color: #4E1D00;
		color: #EADFCE;
		font-weight: bold;
	}
</style>