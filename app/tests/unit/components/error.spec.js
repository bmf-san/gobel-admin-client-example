import { shallowMount } from "@vue/test-utils";
import Error from "@/components/Error.vue";

describe("Error.vue", () => {
  it("matched snapshot", () => {
    const wrapper = shallowMount(Error);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
