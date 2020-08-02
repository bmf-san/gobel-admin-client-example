import { shallowMount } from "@vue/test-utils";
import Loader from "@/components/Loader.vue";

describe("Loader.vue", () => {
  it("init", () => {
    const wrapper = shallowMount(Loader);
    expect(wrapper.isVueInstance).toBeTruthy();
  });

  it("matched snapshot", () => {
    const wrapper = shallowMount(Loader);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
