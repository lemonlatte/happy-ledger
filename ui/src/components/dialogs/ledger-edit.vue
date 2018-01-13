<style lang="scss" scoped>
  .mask {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: rgba(15, 15, 15, 0.7);
    text-align: center;
  }

  .dialog {
    background-color: white;
    border: solid 1px black;
    display: inline-block;
    width: 400px;
    height: 100px;

    .close {
      float: right;
    }
  }
</style>

<template>
  <div v-if="editingLedger.id">
    <div class="mask">
      <div class="dialog">
        <button class="close" @click="close">Close</button>

        <form>
          <p>
            Name
            <input type="text" :placeholder="editingLedger.name" v-model="newName">
          </p>

        </form>
        <button @click="ok">確定</button>
        <button @click="close">取消</button>
      </div>
    </div>
  </div>
</template>
<script>
  export default {
    props: {
      editingLedger: {},
    },

    methods: {
      ok() {
        if (!this.newName) {
          this.$emit('close')
        } else {
          let newName = this.newName
          this.$emit('edit', newName)
        }
        this.newName = ''
      },
      close() {
        this.$emit('close')
        this.newName = ''
      }
    },

    data() {
      return {
        newName: "",
      }
    }
  }
</script>
