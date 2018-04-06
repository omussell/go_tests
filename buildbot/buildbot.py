c['schedulers'].append(schedulers.SingleBranchScheduler(
                            name="testing factory",
                            change_filter=util.ChangeFilter(branch='master'),
                            treeStableTimer=None,
                            #fileIsImportant=config_updated,
                            builderNames=["testingbuilder"]))

c['builders'].append(
    util.BuilderConfig(name="testingbuilder",
      workernames=["buildmaster"],
      factory=testfactory))

testfactory = util.BuildFactory()
testfactory.addStep(steps.ShellCommand(command=["echo", "build"],))

