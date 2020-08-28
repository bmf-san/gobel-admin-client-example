import { shallowMount } from "@vue/test-utils";
import Pagination from "@/components/Pagination.vue";

describe("Pagination.vue", () => {
  it("init", () => {
    const wrapper = shallowMount(Pagination);
    expect(wrapper.isVueInstance).toBeTruthy();
  });

  it("matched snapshot", () => {
    const wrapper = shallowMount(Pagination);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
