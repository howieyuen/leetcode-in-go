package _207_course_schedule

import (
	"testing"
)

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: true,
		},
		{
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
			if got := canFinish1(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish1() = %v, want %v", got, tt.want)
			}
			if got := canFinish2(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish2() = %v, want %v", got, tt.want)
			}
		})
	}
}
