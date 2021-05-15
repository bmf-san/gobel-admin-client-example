import { shallowMount } from "@vue/test-utils";
import Footer from "@/components/Footer.vue";

describe("Footer.vue", () => {
  it("init", () => {
    const wrapper = shallowMount(Footer);
    expect(wrapper.isVueInstance).toBeTruthy();
  });

  it("matched snapshot", () => {
    const wrapper = shallowMount(Footer);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
