c['builders'].append(
    util.BuilderConfig(name="test",
      workernames=["buildmaster"],
      factory=testfactory))

testfactory = util.BuildFactory()
testfactory.addStep(steps.ShellCommand(command=["echo", "build"],))

