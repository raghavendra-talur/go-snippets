func executeCmd(ctx context.Context, cmdArgs []string) (cmdout, cmderr bytes.Buffer, err error) {
	var stdout, stderr bytes.Buffer
	var stdoutWriters, stderrWriters []io.Writer

	cmd := exec.CommandContext(ctx, cmdArgs[0], cmdArgs[1:]...)

	stdoutWriters = append(stdoutWriters, &stdout)
	stderrWriters = append(stderrWriters, &stderr)

	if IsStreaming(ctx) {
		stdoutWriters = append(stdoutWriters, os.Stdout)
		stderrWriters = append(stderrWriters, os.Stderr)
	}

	cmd.Stdout = io.MultiWriter(stdoutWriters...)
	cmd.Stderr = io.MultiWriter(stderrWriters...)

	if err := cmd.Run(); err != nil {
                return stdout, stderr, fmt.Errorf("failed to execute command\n\tcommand: %s:\n\terrstring: %s\texitStatus: %w", strings.Join(cmdArgs, " "), stderr.String(), err)
	}

	return stdout, stderr, nil
}
