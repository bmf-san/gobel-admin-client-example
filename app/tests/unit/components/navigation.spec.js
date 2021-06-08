import { shallowMount } from "@vue/test-utils";
import Error from "@/components/Error.vue";

describe("Error.vue", () => {
  it("init", () => {
    const $route = {
      path: "/some/path"
    };
    const wrapper = shallowMount(Error, {
      mock: {
        $route
      }
    });
    expect(wrapper.isVueInstance).toBeTruthy();
  });

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
