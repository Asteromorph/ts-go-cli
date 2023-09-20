import { Operation } from "../config";
import Projector from "../projector";

function createData() {
  return {
    projector: {
      "/": {
        "foo": "bar1",
        "fem": "test"
      },
      "/foo": {
        "foo": "bar2",
      },
      "/foo/bar": {
        "foo": "bar3",
      },
    }
  }
}

function getProjector(pwd: string, data = createData()): Projector {
  return new Projector({
    args: [],
    operation: Operation.Print,
    pwd,
    config: "Hello"
  }, data)
}

test("getAllValues", function() {
  const proj = getProjector("/foo/bar")
  expect(proj.getAllValues()).toEqual({
    "fem": "test",
    "foo": "bar3"
  })
})

test("getValue", function() {
  let proj = getProjector("/foo/bar")
  expect(proj.getValue("foo")).toEqual(
     "bar3"
  );
  proj = getProjector("/foo")
  expect(proj.getValue("foo")).toEqual(
     "bar2"
  );
  expect(proj.getValue("fem")).toEqual(
     "test"
  );
})

test("setValue", function() {
  let data = createData();
  let proj = getProjector("/foo/bar", data)
  proj.setValue("foo", "baz")
  expect(proj.getValue("foo")).toEqual(
    "baz"
  )
  proj.setValue("fem", "test2")
  expect(proj.getValue("fem")).toEqual(
    "test2"
  )

  proj = getProjector("/", data)
  expect(proj.getValue("fem")).toEqual(
    "test"
  )
})

test("removeValue", function() {
  const proj = getProjector("/foo/bar")
  proj.removeValue("fem")
  expect(proj.getValue("fem")).toEqual(
    "test"
  )

  proj.removeValue("foo")
  expect(proj.getValue("foo")).toEqual(
    "bar2"
  )
})
