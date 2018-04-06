testfactory = util.BuildFactory()
testfactory.addStep(steps.ShellCommand(command=["echo", "build"],))

c['builders'].append(
    util.BuilderConfig(name="testingbuilder",
      workernames=["buildmaster"],
      factory=testfactory))

c['schedulers'].append(schedulers.SingleBranchScheduler(
                            name="testing factory",
                            change_filter=util.ChangeFilter(branch='master'),
                            treeStableTimer=None,
                            #fileIsImportant=config_updated,
                            builderNames=["testingbuilder"]))
