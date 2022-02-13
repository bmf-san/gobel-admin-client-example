import { shallowMount } from "@vue/test-utils";
import Error from "@/components/Error.vue";

describe("Error.vue", () => {
  it("matched snapshot", () => {
    const $route = {
      path: "/some/path"
    };
    const wrapper = shallowMount(Error, {
      mock: {
        $route
      }
    });
    expect(wrapper.html()).toMatchSnapshot();
  });
});
