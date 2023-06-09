import { shallowMount } from "@vue/test-utils";
import Alert from "@/components/atoms/Alert";

describe("Alert", () => {
  let wrapper;
  beforeEach(() => {
    wrapper = shallowMount(Alert);
  });

  it("matches snapshot", () => {
    expect(wrapper.html()).toMatchSnapshot();
  });
});
